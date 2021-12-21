console.log("Current time: " + new Date())
import *
        as
        THREE
    from
        'https://cdn.skypack.dev/three@v0.135.0';



const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);
const renderer = new THREE.WebGLRenderer();
renderer.setSize(window.innerWidth, window.innerHeight);
document.body.appendChild(renderer.domElement);

const geometry = new THREE.BoxGeometry();
const material = new THREE.MeshBasicMaterial({color: 0x00ff00});
const cube = new THREE.Mesh(geometry, material);
scene.add(cube);
camera.position.z = 5;
camera.position.y = .1;
camera.position.x = 5;

const plane = new THREE.Plane(new THREE.Vector3(0, 1, 0), 0);
const helper = new THREE.PlaneHelper(plane, 100, 0xffff00);
scene.add(helper);

camera.position.x = 30
camera.position.y = 30
camera.position.z = 30

camera.lookAt(0, 0, 0)


function animate() {
    requestAnimationFrame(animate);
    renderer.render(scene, camera);
}


    protobuf.load("messages.proto", function (err, root) {
        if (err)
            throw err;

        let FromServer = root.lookupType("rts.FromServer")
        let FromClient = root.lookupType("rts.FromClient")

        let greetings = {hello: {greeting: "Hello -- I'm a web browser"}}
        FromClient.setup()
        let errMsg = FromClient.verify(greetings)
        if (errMsg) {
            throw Error(errMsg)
        }
        let sendThis = FromClient.create(greetings);
        let sendThisBuffer = FromClient.encode(sendThis).finish();

        let webSocket = new WebSocket("ws://localhost:8000/ws");
        webSocket.binaryType = 'arraybuffer';

        webSocket.onclose = function (event) {
            console.log("Closed it " + event)
        }
        webSocket.onopen = function (event) {
            console.log("websocket open " + JSON.stringify(event))
            webSocket.send(sendThisBuffer)
        }
        webSocket.onmessage = function (event) {
            let msg = event.data
            let message = FromServer.decode(new Uint8Array(msg))
            console.log("Do a message? " + JSON.stringify(message))
            console.log("can find the stuff: " + message.hello.greeting)
        }
    })

animate()