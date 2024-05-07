fetch("http://127.0.0.1:33212/startUniProxy", {
  method: "POST",
  headers: {
    "Content-Type": "application/json"
  },
  body: JSON.stringify({
    tag: "shadowsocks_1",
    uuid: "72493186-abeb-479d-982c-b7dd7a0afc6d",
    global_mode: false
  })
})
.then(response => response.json())
.then(data => console.log('Success:', data))
.catch(error => console.error('Error:', error));