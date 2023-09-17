import { commands } from "./commands";

let socket = new WebSocket("ws://127.0.0.1:8080/ws");
console.log("Attempting Connection...");
socket.onopen = () => {
  console.log("Successfully Connected");
};

let term;

socket.onmessage = (event) => {
  console.log("received from server", event.data);
  term.echo(event.data);
  term.resume();
};
socket.onclose = (event) => {
  console.log("Socket Closed Connection: ", event);
  socket.send("Client Closed!");
};

socket.onerror = (error) => {
  console.log("Socket Error: ", error);
};

// short hand for document .ready()
// need it because I have a src file in head, so this should only be called after document is ready
$(function () {
  term = $("#terminal").terminal(
    (command, term) => {
      term.pause();
      socket.send(command);
    },
    {
      prompt: "goshell% ",
      greetings: "Welcome to your shell!"
    }
  );
});

const inputEl = document.getElementById("learn-input")
inputEl.addEventListener('keydown', (event) => {
  if (event.key == 'Enter') {
    event.preventDefault()
    const textDiv = document.getElementById("text-display")
    const cmd = event.target.value
    let text = commands[cmd]

    if (!text) {
      text = "Command not found"
    }
    textDiv.innerHTML = `
    <p class="cmd-info">
    ${text}
    </p>
    `
  }
})