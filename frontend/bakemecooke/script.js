console.log("saaas")


const ip = "192.168.108.194"
const adress = "http://"+ip;//127.0.0.1";
const wsAdress = "ws://"+ip//127.0.0.1";
const port = ":7070";
const gameRoute = "/game";
const connectRoute = "/connect"




function first(){
    const request = new XMLHttpRequest();
    request.open("GET", adress+port+gameRoute, true);
    //request.setRequestHeader("Access-Control-Request-Headers", "*");
    //request.setRequestHeader("Access-Control-Allow-Origin", "*");
    //request.setRequestHeader("Access-Control-Allow-Headers", "*");
    request.setRequestHeader("Content-Type", "application/json");
    request.setRequestHeader("Authorization", "Bearer your_token");
    request.setRequestHeader("Custom-Header", "Custom Value");
    request.responseType = "text";
    request.onload = () => {
        if (request.readyState == 4 && request.status == 200) {
            const data = request.response;
            console.log(data);
        } else {
            console.log(`Error: ${request.status}`);
        }
    };
    request.send();
}

async function second() {
    const tt = adress+port+gameRoute;    
    console.log(
        fetch(tt)
            .then(response =>response.text().then(val=> val)));
}

async function third(){

    setInterval(second,1000);
}

//third();
const messageFieldId= "messagefield";



const socket = new WebSocket(wsAdress+port+connectRoute);
console.log("sock",socket);
// Обработчик события открытия соединения
socket.onopen = function(event) {
    console.log("Соединение установлено");
    
    // Отправка данных на сервер
    socket.send("Привет, сервер!");
  };
  
  // Обработчик события получения сообщения от сервера
  socket.onmessage = function(event) {
    const messageFiled = document.getElementById("messagefield");
    console.log("Получено сообщение от сервера:", event.data);
    console.log(messageFiled);
    messageFiled.innerText += event.data+"\n";
  };
  
  // Обработчик события закрытия соединения
  socket.onclose = function(event) {
    console.log("Соединение закрыто");
  };

function getText() {
    var inputElement = document.getElementById("myInput");
    var text = inputElement.value;
    console.log(socket);
    socket.send(text);
}

/*

    add_header 'Access-Control-Allow-Origin' $http_origin;
    add_header 'Access-Control-Allow-Headers' 'x-requested-with, Content-Type, origin, authorization, accept, client-security-token';
    add_header 'Access-Control-Allow-Methods' 'GET,POST,OPTIONS,PUT,PATCH,DELETE';
    add_header 'Access-Control-Allow-Credentials' 'true';

*/