import React, { CSSProperties, ChangeEvent, useEffect, useState } from 'react';
import logo from './logo.svg';
import './App.css';
import { io, Socket } from 'socket.io-client';

export interface StylesDictionary {
  [Key: string]: CSSProperties
}

const suckit = new WebSocket("ws://192.168.108.194:7070/connect");
function App() {
  
  const [inputValue, setInputValue] = useState('');
  const [messages, setMessages] = useState<string[]>([]);

  function addMessage(str: string){
    //const newMessages = messages;
    //newMessages.push(str);
    //console.log(newMessages);
    //  setMessages(/*newMessages*/[ ...messages,str ]);
    setMessages((ms)=>[ ...ms,str ]);
  
  }
  useEffect(()=>{
    console.log("use effect updated");
    suckit.onclose = (event: CloseEvent)=>{
      console.log("Соединение закрыто");
    };
    suckit.onmessage = (event: MessageEvent<string>)=>{
      console.log("Получено сообщение от сервера:", event.data);
      addMessage(event.data);
    };
    suckit.onopen = (event: Event)=>{
      console.log("Соединение установлено");
    
      // Отправка данных на сервер
      suckit.send("Привет, сервер!");
    };
  },[]);

  function sendText(){
      
      suckit.send(inputValue); 
      console.log("Бляьт на кнопку нажал"); 
  }

  function getValue(event: ChangeEvent<HTMLInputElement>){
    const val = event.target.value;
    setInputValue(val);
  }

  return (
    <div className="App">
      <div className='cube'></div>
      <input type="text" value={inputValue} onChange={getValue} id="myInput"/>
      <button value={inputValue} onClick={sendText}>Получить текст</button>
      <div id="messagefield">
        {messages.map((val, index)=><div key={index}>{val}</div>)}
      </div>
    </div>
  );
}

export default App;
/*


ing&t=OkKx8d1 net::ERR_FAILED 404 (Not Found)
xhrSendProcessor @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3751
send @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3767
create @ polling.js:298
Request @ polling.js:237
request @ polling.js:190
doPoll @ polling.js:215
poll @ polling.js:96
doOpen @ polling.js:56
open @ transport.js:46
open @ socket.js:170
Socket @ socket.js:111
open @ manager.js:108
(anonymous) @ manager.js:328
setTimeout (async)
reconnect @ manager.js:321
(anonymous) @ manager.js:331
onError @ manager.js:123
Emitter.emit @ index.mjs:136
onError @ socket.js:541
Emitter.emit @ index.mjs:136
onError @ transport.js:38
(anonymous) @ polling.js:218
Emitter.emit @ index.mjs:136
onError @ polling.js:320
(anonymous) @ polling.js:294
setTimeout (async)
xhr.onreadystatechange @ polling.js:293
XMLHttpRequest.send (async)
xhrSendProcessor @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3751
send @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3767
create @ polling.js:298
Request @ polling.js:237
request @ polling.js:190
doPoll @ polling.js:215
poll @ polling.js:96
doOpen @ polling.js:56
open @ transport.js:46
open @ socket.js:170
Socket @ socket.js:111
open @ manager.js:108
(anonymous) @ manager.js:328
setTimeout (async)
reconnect @ manager.js:321
(anonymous) @ manager.js:331
onError @ manager.js:123
Emitter.emit @ index.mjs:136
onError @ socket.js:541
Emitter.emit @ index.mjs:136
onError @ transport.js:38
(anonymous) @ polling.js:218
Emitter.emit @ index.mjs:136
onError @ polling.js:320
(anonymous) @ polling.js:294
setTimeout (async)
xhr.onreadystatechange @ polling.js:293
XMLHttpRequest.send (async)
xhrSendProcessor @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3751
send @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3767
create @ polling.js:298
Request @ polling.js:237
request @ polling.js:190
doPoll @ polling.js:215
poll @ polling.js:96
doOpen @ polling.js:56
open @ transport.js:46
open @ socket.js:170
Socket @ socket.js:111
open @ manager.js:108
(anonymous) @ manager.js:328
setTimeout (async)
reconnect @ manager.js:321
(anonymous) @ manager.js:331
onError @ manager.js:123
Emitter.emit @ index.mjs:136
onError @ socket.js:541
Emitter.emit @ index.mjs:136
onError @ transport.js:38
(anonymous) @ polling.js:218
Emitter.emit @ index.mjs:136
onError @ polling.js:320
(anonymous) @ polling.js:294
setTimeout (async)
xhr.onreadystatechange @ polling.js:293
XMLHttpRequest.send (async)
xhrSendProcessor @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3751
send @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3767
create @ polling.js:298
Request @ polling.js:237
request @ polling.js:190
doPoll @ polling.js:215
poll @ polling.js:96
doOpen @ polling.js:56
open @ transport.js:46
open @ socket.js:170
Socket @ socket.js:111
open @ manager.js:108
(anonymous) @ manager.js:328
setTimeout (async)
reconnect @ manager.js:321
(anonymous) @ manager.js:331
onError @ manager.js:123
Emitter.emit @ index.mjs:136
onError @ socket.js:541
Emitter.emit @ index.mjs:136
onError @ transport.js:38
(anonymous) @ polling.js:218
Emitter.emit @ index.mjs:136
onError @ polling.js:320
(anonymous) @ polling.js:294
setTimeout (async)
xhr.onreadystatechange @ polling.js:293
XMLHttpRequest.send (async)
xhrSendProcessor @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3751
send @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3767
create @ polling.js:298
Request @ polling.js:237
request @ polling.js:190
doPoll @ polling.js:215
poll @ polling.js:96
doOpen @ polling.js:56
open @ transport.js:46
open @ socket.js:170
Socket @ socket.js:111
open @ manager.js:108
(anonymous) @ manager.js:328
setTimeout (async)
reconnect @ manager.js:321
(anonymous) @ manager.js:331
onError @ manager.js:123
Emitter.emit @ index.mjs:136
onError @ socket.js:541
Emitter.emit @ index.mjs:136
onError @ transport.js:38
(anonymous) @ polling.js:218
Emitter.emit @ index.mjs:136
onError @ polling.js:320
(anonymous) @ polling.js:294
setTimeout (async)
xhr.onreadystatechange @ polling.js:293
XMLHttpRequest.send (async)
xhrSendProcessor @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3751
send @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3767
create @ polling.js:298
Request @ polling.js:237
request @ polling.js:190
doPoll @ polling.js:215
poll @ polling.js:96
doOpen @ polling.js:56
open @ transport.js:46
open @ socket.js:170
Socket @ socket.js:111
open @ manager.js:108
(anonymous) @ manager.js:328
setTimeout (async)
reconnect @ manager.js:321
(anonymous) @ manager.js:331
onError @ manager.js:123
Emitter.emit @ index.mjs:136
onError @ socket.js:541
Emitter.emit @ index.mjs:136
onError @ transport.js:38
(anonymous) @ polling.js:218
Emitter.emit @ index.mjs:136
onError @ polling.js:320
(anonymous) @ polling.js:294
setTimeout (async)
xhr.onreadystatechange @ polling.js:293
XMLHttpRequest.send (async)
xhrSendProcessor @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3751
send @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3767
create @ polling.js:298
Request @ polling.js:237
request @ polling.js:190
doPoll @ polling.js:215
poll @ polling.js:96
doOpen @ polling.js:56
open @ transport.js:46
open @ socket.js:170
Socket @ socket.js:111
open @ manager.js:108
(anonymous) @ manager.js:328
setTimeout (async)
reconnect @ manager.js:321
(anonymous) @ manager.js:331
onError @ manager.js:123
Emitter.emit @ index.mjs:136
onError @ socket.js:541
Emitter.emit @ index.mjs:136
onError @ transport.js:38
(anonymous) @ polling.js:218
Emitter.emit @ index.mjs:136
onError @ polling.js:320
(anonymous) @ polling.js:294
setTimeout (async)
xhr.onreadystatechange @ polling.js:293
XMLHttpRequest.send (async)
xhrSendProcessor @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3751
send @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3767
create @ polling.js:298
Request @ polling.js:237
request @ polling.js:190
doPoll @ polling.js:215
poll @ polling.js:96
doOpen @ polling.js:56
open @ transport.js:46
open @ socket.js:170
Socket @ socket.js:111
open @ manager.js:108
(anonymous) @ manager.js:328
setTimeout (async)
reconnect @ manager.js:321
(anonymous) @ manager.js:331
onError @ manager.js:123
Emitter.emit @ index.mjs:136
onError @ socket.js:541
Emitter.emit @ index.mjs:136
onError @ transport.js:38
(anonymous) @ polling.js:218
Emitter.emit @ index.mjs:136
onError @ polling.js:320
(anonymous) @ polling.js:294
setTimeout (async)
xhr.onreadystatechange @ polling.js:293
XMLHttpRequest.send (async)
xhrSendProcessor @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3751
send @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3767
create @ polling.js:298
Request @ polling.js:237
request @ polling.js:190
doPoll @ polling.js:215
poll @ polling.js:96
doOpen @ polling.js:56
open @ transport.js:46
open @ socket.js:170
Socket @ socket.js:111
open @ manager.js:108
(anonymous) @ manager.js:328
setTimeout (async)
reconnect @ manager.js:321
(anonymous) @ manager.js:331
onError @ manager.js:123
Emitter.emit @ index.mjs:136
onError @ socket.js:541
Emitter.emit @ index.mjs:136
onError @ transport.js:38
(anonymous) @ polling.js:218
Emitter.emit @ index.mjs:136
onError @ polling.js:320
(anonymous) @ polling.js:294
setTimeout (async)
xhr.onreadystatechange @ polling.js:293
XMLHttpRequest.send (async)
xhrSendProcessor @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3751
send @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3767
create @ polling.js:298
Request @ polling.js:237
request @ polling.js:190
doPoll @ polling.js:215
poll @ polling.js:96
doOpen @ polling.js:56
open @ transport.js:46
open @ socket.js:170
Socket @ socket.js:111
open @ manager.js:108
(anonymous) @ manager.js:328
setTimeout (async)
reconnect @ manager.js:321
(anonymous) @ manager.js:331
onError @ manager.js:123
Emitter.emit @ index.mjs:136
onError @ socket.js:541
Emitter.emit @ index.mjs:136
onError @ transport.js:38
(anonymous) @ polling.js:218
Emitter.emit @ index.mjs:136
onError @ polling.js:320
(anonymous) @ polling.js:294
setTimeout (async)
xhr.onreadystatechange @ polling.js:293
XMLHttpRequest.send (async)
xhrSendProcessor @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3751
send @ main.js?attr=MSdVUVOWkVuy1CUTx1GLWAZJRki0LJLv6w1zo4R_trEpsH1sULSvca7r-WChrn7V:3767
create @ polling.js:298
Request @ polling.js:237
request @ polling.js:190
doPoll @ polling.js:215
poll @ polling.js:96
doOpen @ polling.js:56
open @ transport.js:46
open @ socket.js:170
Socket @ socket.js:111
open @ manager.js:108
(anonymous) @ manager.js:328
setTimeout (async)
reconnect @ manager.js:321
(anonymous) @ manager.js:331
onError @ manager.js:123
Emitter.emit @ index.mjs:136
onError @ socket.js:541
Emitter.emit @ index.mjs:136
onError @ transport.js:38
(anonymous) @ polling.js:218
Emitter.emit @ index.mjs:136
onError @ polling.js:320
(anonymous) @ polling.js:294
setTimeout (async)
xhr.onreadystatechange @ polling.js:293
localhost/:1 Access to XMLHttpRequest at '' from origin '' has been blocked by CORS policy: No 'Access-Control-Allow-Origin' header is present on the requested resource.

*/