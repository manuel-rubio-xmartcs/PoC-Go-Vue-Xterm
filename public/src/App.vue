<template>
  <div>
    <h1>Hola</h1>
    <v-btn @click="connect">Conectar</v-btn>
    <div id="xterminal"></div>
    <input type="text" class="terminal xterm xterm-dom-renderer-owner-1" style="display:none;" id="xter" />
  </div>
</template>

<script>
  export default {
    methods: {
      connect() {
        console.log("Trying to connect");
        var term = new Terminal();

        let connection = new WebSocket("ws://localhost:8080/ws");
        
        connection.onmessage = (event) => {
          console.log(event.data);
          term.write(event.data+"\n\r");
          document.getElementById("xter").value = "";
        }
        
        document.getElementById("xter").addEventListener("keyup", function(event) {
          event.preventDefault();
          if (event.key === 'Enter') {
            term.write("$> "+ document.getElementById("xter").value + "\n\r")
            connection.send(document.getElementById("xter").value)
          }
        });
        
        connection.onopen = () => {
          console.log("Connected!")
          term.open(document.getElementById("xterminal"))
          document.getElementById("xter").style = "display: block; font-family: monospace; background-color: black; width: 100%;"
        }


        connection.onerror = (event) => {
          console.log("Connection refused!")
          console.log(event)
        }
      },
    },
  };

  
</script>

