let sseUri = "http://localhost:1235/sse" //"http://localhost:8090/sse"
let source = new EventSource(sseUri);
let live = document.querySelector('.badge')

source.onopen = function(){
  live = document.querySelector('.badge')
  live.style.backgroundColor = '#ff334b';
}

source.onerror = function() {
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
    console.log(event)
    var message = event.data
    console.log(message)

    const obj = JSON.parse(message)
    const Details = document.createElement("details");
    const Summary = document.createElement("summary");
    const Content = document.createElement("p");
    const Badge = document.createElement("span");
    const Attachment = document.createElement(null);

    Badge.className = "chip"
    incomingMsg = obj.MessageText.replace(/\"/g, "")
    // const result = 10 > 5 ? 'yes' : 'no';
    const from = obj.Sender === obj.Group ? obj.Sender : obj.Sender + " / " + obj.Group;
    const msg = obj.MessageCaption === "" ? 
    incomingMsg : 
    incomingMsg +'<br>'+ "Caption: " + obj.MessageCaption;

    const body = obj.Uri === "" ? msg : msg + '<br>' + "<a href='" + obj.Uri + "' target='_blank'>Open attachment</a>"

    Details.setAttribute("data-type", obj.MessageType)
    Details.setAttribute("data-url", obj.Uri)
    Details.setAttribute("data-id", obj.MessageID)
    Details.setAttribute("data-sender", obj.Sender)

    Content.innerHTML = `Text: <span style="color:green">${body}</span>` + "<br><label for='reply'>Reply: </label>" +
    "<input type='text' id='reply' name='reply' style='width:80%'>" +
   // "<input type='button' onclick='location.href=`https://google.com`;' value='WhatsApp Reply' />"
   "<input type='button' onclick='sendwa(this);' value='WhatsApp Reply' />"

    if(obj.MessageType === "image") {
      Attachment.innerHTML = '<br><br>' + "<img src='"+obj.Uri+"' alt='Message attachment' width: 80%;' height='600'>"
    } else if(obj.MessageType === "audio"){
    //  console.log(`Audio source: ${obj.Uri}`)
      Attachment.innerHTML = '<br>' + "<audio controls> "+
        "<source src='"+obj.Uri+"' type='audio/ogg'>" +
        "<source src='"+obj.Uri+"' type='audio/mpeg'>" +
        "<source src='"+obj.Uri+"' type='audio/oga'>" +
     //   "<source src='test.oga' type='audio/ogg; codecs=`vorbis`'></source>" +
        "Your browser does not support the audio element." +
      "</audio>"
    } else if(obj.MessageType === "video"){
      Attachment.innerHTML = '<br>' + "<video width: 100%;' height='600' controls> " +
        "<source src='"+obj.Uri+"' type='video/ogg'>" +
        "<source src='"+obj.Uri+"' type='video/mp4'>" +
        "<source src='"+obj.Uri+"' type='video/m4v'>" +
        "Your browser does not support the video tag." +
      "</video>"
    } else if(obj.MessageType === "document" && obj.Uri.split('.').pop() === 'pdf'){ //; obj.Uri.type === 'application/pdf'){
      Attachment.innerHTML = '<br>' +  "<embed id='pdf' type='application/pdf'" +
      "src='"+obj.Uri+"' style='width: 100%;' height='600'>"
    } else if (obj.MessageType === "document" && obj.Uri.split('.').pop() !== 'pdf'){
      var ext = obj.Uri.split('.').pop() // obj.Uri.split('.').reverse()[0]
      Content.innerHTML += `<br> file recieved of of extension ${ext}`
    }

    console.log(body)
    let isMentioned = body.includes("@966506889946") || body.includes("966506889946@");
    console.log(`is mentioned: ${isMentioned}`)
    if (isMentioned){
      console.log(`Mentioned`)
      Badge.classList.add('danger')
      Badge.textContent = "Mentioned" // "<span class='chip danger'>Mentioned</span>"
      console.log(`Badge is: ${Badge.innerHTML}`)
    }

    Summary.innerHTML = "From: "+ '<code>' + obj.Name + '</code>' + " Number: " + from + " By: " + obj.Time + " "
    // "<span class='chip danger'>Mentioned</span>"
    
    Summary.appendChild(Badge);
    Details.appendChild(Summary);

    Content.appendChild(Attachment);
    Details.appendChild(Content);

    Details.addEventListener('click', function(){
      chip = this.querySelector('.chip')
      if (chip != null) {
        chip.remove();
      }
    })

  //  let container = document.querySelector(".container");

  //  Details.after(container);
    document.body.insertBefore(Details, document.body.children[3]);

  //  console.log("0:",document.body.children[0]) // head
  //  console.log("1:",document.body.children[1]) // script
  //   console.log("2:",document.body.children[2]) // container

   let ds= document.querySelectorAll('details');
   window.addEventListener('click', function(event){
      ds.forEach(function(d){
        const isClickInside = d.contains(event.target);
        if(!isClickInside){
          d.removeAttribute('open')
        }
      })
    }, true);

}


function sendwa(e){
  sender = e.parentNode.parentNode.dataset.sender; // .querySelector("details")
  reply = document.querySelector("#reply");
  replyText = reply.value

  var today = new Date();
  var date = today.getFullYear()+'-'+(today.getMonth()+1)+'-'+today.getDate();
  var time = today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds();
  var dateTime = date+' '+time;

  const r = document.createElement("p");
  r.innerHTML += `[${dateTime}] I replied: <span style="color:blue">${replyText}</span>`
  // e.parentNode.parentNode.insertBefore(r, e.parentNode);

 // res = window.send(sender, replyText)
 // res.then(data){alert(data)}
  //alert(res)

 // res = window.send(sender, replyText)
 // alert(res)
 // res1 = res.json()
 // alert(res1)
 /* res.then(
    (res => res.json())
    .then(data => {
      alert(data)
      // enter you logic when the fetch is successful
      reply.value = ""
      e.parentNode.parentNode.appendChild(r)
     })
     .catch(error => {
     // enter your logic for when there is an error (ex. error toast)
      alert(data.message)
   //  alert(error)
    })
  ); */
 // res = res.json()
 // alert(res)
 /* .then(data => {
    alert(data)
    // enter you logic when the fetch is successful
    reply.value = ""
    e.parentNode.parentNode.appendChild(r)
   })
   .catch(error => {
    // enter your logic for when there is an error (ex. error toast)
     alert(data.message)
   }) */

 // reply.value = ""
    fetch('http://localhost:1235/reply', {
      method: 'POSTS',
      headers: {
        'Content-Type': 'application/json',
      },
       body: JSON.stringify({
         // your expected POST request payload goes here
         sender: sender,
         message: replyText
          })
    }) 
    .then(res => res.json())
    .then(data => {
       // enter you logic when the fetch is successful
       reply.value = ""
       e.parentNode.parentNode.appendChild(r)
      })
      .catch(error => {
      // enter your logic for when there is an error (ex. error toast)
       alert(data.message)
    //  alert(error)
     }) 

     

 // alert(sender+" "+reply.value)
 // window.send(sender, reply.value)
  //reply.innerHTML = ""
//  window.send(sender, reply.value)
//  response = encodeURIComponent(reply.value)
//  console.log(e.parentNode.parentNode)
//  console.log(sender)
}
/*
  window.open(
    "https://wa.me/" + sender + "/?text=" + response,
    '_blank' // <- This is what makes it open in a new window.
  );

// Below are alternate options
 // location.href = r
//  let a= document.createElement('a');
//  a.target= '_blank';
//  a.href= "https://wa.me/" + sender + "/?text=" + response;
//  a.click();

  // <a id="anchorID" href="mynewurl" target="_blank"></a>
}  */