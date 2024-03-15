const express = require('express')
const {join} = require('node:path')
const http = require('http'); 
const { Server } = require('socket.io');
const cors = require('cors'); // Import the CORS middleware

const app = express()
const port = 3000


const server = http.createServer(app)

const io = new Server(server);
// var io = require('socket.io').listen(server);
app.use(cors()); // Enable CORS for all routes
app.get('/', (req, res) => {
  res.sendFile(join(__dirname, 'index.html'))
})

io.on('connection', (socket) => {

  socket.on('chat message', (msg) => {
    io.emit('chat message', msg);
  });

  socket.emit('test emit')

  socket.on('world', (arg)=>{
    console.log(arg)
  })

  socket.emit('hello', 'world')


});

server.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})

