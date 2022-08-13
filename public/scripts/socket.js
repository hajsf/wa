let sseUri = "http://localhost:1235/sse" //"http://localhost:8090/sse"
let source = new EventSource(sseUri);
let live = document.querySelector('.badge')

source.onopen = function(){
  live = document.querySelector('.badge')
  live.style.backgroundColor = '#ff334b';
}

source.onerror = function() {
  document.querySelector('#qr').innerHTML = "";
  document.querySelector('#message').innerHTML = "Server is down at the host machine ";
  live.style.backgroundColor = '#30000614';
  setInterval(() => {
    if (source.readyState == EventSource.CLOSED) {
        live.style.backgroundColor = '#30000614';
        source.close();
        source = new EventSource("http://localhost:1235/sse");
    } 
}, 3000);
}

source.onmessage = function (event) {
  document.querySelector('#qr').innerHTML = "";
  const obj = JSON.parse(event.data);
  document.querySelector('#message').innerHTML = `Message ${obj.MessageText} received from ${obj.SenderName}`
}

// will be called automatically whenever the server sends a message with the event field set to "qr"
// echo "event: notification\ndata: {"time": "' . $curDate . '"}';
// fmt.Fprint(w, "event: notification\ndata: %v\n\n", c) then followed by fmt.Fprintf(w, "data: %v\n\n", c)
source.addEventListener("notification", function(event) {
  console.log("notification recieved")
 // console.log(event.data)
  document.querySelector('#qr').innerHTML = "";
  var message = event.data
  document.querySelector('#message').innerHTML = message;
});

source.addEventListener("result", function(event) {
  console.log("result recieved")
  // console.log(event.data)
   var message = event.data
   document.querySelector('#message').innerHTML = message;
 });

 source.addEventListener("granted", function(event) {
    location.replace('http://localhost:1235//messages');
 });

source.addEventListener("qrCode", function(event) {
  var message = event.data
  console.log(event.data)
  document.querySelector('#qr').innerHTML = "";
  document.querySelector('#message').innerHTML = "Scan the QR code from the WhatsApp application";
  var qrcode = new QRCode("qr", {
    text: message,
    width: 128,
    height: 128,
    colorDark : "#000000",
    colorLight : "#ffffff",
    correctLevel : QRCode.CorrectLevel.M
});
});