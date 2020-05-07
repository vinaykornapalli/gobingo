const evtSource = new EventSource("http://localhost:8080/event?id=1234");

evtSource.onmessage = function(event) {
    const newElement = document.createElement("li");
    const eventList = document.getElementById("list");
    console.log(event.data)
    newElement.innerHTML = "message: " + event.data;
    eventList.appendChild(newElement);
}