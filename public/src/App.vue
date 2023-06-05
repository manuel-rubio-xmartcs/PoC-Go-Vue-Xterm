<template>
  <div>
    <h1>Hola</h1>
    <v-btn @click="connect">Conectar</v-btn>
    <div id="xterminal"></div>
  </div>
</template>

<script>
  export default {
    methods: {
      connect() {
        let command = "";
        let socket = new WebSocket("ws://localhost:8080/ws");
        console.log("Connected to websocket!")
        let terminal = new Terminal();

        function clearTerm(command) {
          var inputLengh = command.length;
          for (var i = 0; i < inputLengh; i++) {
            terminal.write('\b \b');
          }
        }

        function prompt(term) {
          console.log("prompt")
          command = '';
          terminal.write('\r\n$ ');
        }

        socket.onmessage = (event) => {
          console.log("msg")
          terminal.write(event.data);
        }

        function runCommand(terminal, command) {
          if (command === "clear" || command === "cls") {
            clearTerm(command);
            terminal.clear();
            return
          }
          if (command.length > 0) {
            console.log("runCommand")
            socket.send(command + '\n');
            terminal.write("\n\r")
          }
        }


        terminal.open(document.getElementById("xterminal"))

        terminal.onData(e => {
          console.log("onData:", e)
          switch (e) {
            case '\u0003': // Ctrl+C
                terminal.write('^C');
                prompt(terminal);
                break;
            case '\r': // Enter
                runCommand(terminal, command);
                command = '';
                break;
            case '\u007F': // Backspace (DEL)
                // Do not delete the prompt
                if (terminal._core.buffer.x > 2) {
                    terminal.write('\b \b');
                    if (command.length > 0) {
                        command = command.substr(0, command.length - 1);
                    }
                }
                break;
            case '\u0009':
                console.log('tabbed', output, ["dd", "ls"]);
                break;
            default:
                if (e >= String.fromCharCode(0x20) && e <= String.fromCharCode(0x7E) || e >= '\u00a0') {
                    command += e;
                    terminal.write(e);
                }
          }
        })
      },
    },

    window:onload = function() {

    }
  };

  
</script>

