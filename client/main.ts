import { serve } from "https://deno.land/std@0.106.0/http/server.ts";
import {
  WebSocketClient,
  StandardWebSocketClient,
  WebSocketError,
} from "https://deno.land/x/websocket@v0.1.3/mod.ts";

console.log("listening on http://localhost:80");
//websocketサーバー
const endpoint = "ws://server:80/ws";

const ws: WebSocketClient = new StandardWebSocketClient(endpoint);
//connection open event
ws.on("open", function () {
  console.log("ws connected!");
  console.log("client: ping!");
  ws.send("ping!");
});
//receive message event
ws.on("message", function (message: any) {
  console.log(`server: ${message.data}`);
  setTimeout(()=> {
    ws.send("ping");
    console.log("client: ping!");
  },1000);
});
//connection close event
ws.on("close", function(){
  console.log("connection closed");
  return
})
//error ocured event
ws.on("error", function (e: WebSocketError) {
  console.log(e);
});

