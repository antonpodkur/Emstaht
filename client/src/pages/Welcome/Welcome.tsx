import type { Component } from "solid-js";
import './welcome.css'
import {createSignal, onMount} from "solid-js";
import * as THREE from "three";
import {LoadGltf} from "../../components/backgrounds/Loader";
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls';


const Welcome : Component = () => {

    let sceneRef: HTMLDivElement | undefined;
    const [sizes, setSizes] = createSignal({width: 600, height: 600})

    onMount(async () => {
        if(sceneRef) {
            setSizes({width: sceneRef.clientWidth, height: sceneRef.clientHeight})
            console.log(sizes())

            // Scene
            const scene = new THREE.Scene()

            // Camera
            const camera = new THREE.PerspectiveCamera(45, sizes().width / sizes().height, 0.1, 100)
            camera.position.y = 0
            camera.position.z = 20
            scene.add(camera)



            // Renderer
            const renderer = new THREE.WebGLRenderer({antialias: true})
            renderer.setSize(sizes().width, sizes().height)
            sceneRef.appendChild(renderer.domElement)
            renderer.pixelRatio = 2

            const controls = new OrbitControls(camera, renderer.domElement)
            controls.autoRotate = true
            controls.enableZoom = false
            controls.update()

            // Totoro
            const totoroMesh = await LoadGltf('Totoro');
            totoroMesh.scale.set(0.5, 0.5, 0.5)
            totoroMesh.position.y = -7

            scene.add(totoroMesh)

            // Point light 1
            const pointLight = new THREE.PointLight('#93ccf6', 1.5, 100)
            pointLight.position.set(3, 10, 10)
            scene.add(pointLight)

            // Point light 2
            const pointLight2 = new THREE.PointLight('#93ccf6', 0.5, 100)
            pointLight2.position.set(2, 5, -10)
            scene.add(pointLight2)

            // Point light 3
            const pointLight3 = new THREE.PointLight('#93ccf6', 0.5, 100)
            pointLight3.position.set(-6, -10, -10)
            scene.add(pointLight3)

            // Ambient light
            const light = new THREE.AmbientLight( 0x404040 ); // soft white light
            scene.add( light );

            renderer.render(scene, camera)

            // Resize handler event listener
            const resizeHandler = (event: Event) => {
                // Update sizes
                setSizes({width: sceneRef!.clientWidth, height: sceneRef!.clientHeight})

                // Update camera aspect and renderer size
                camera.aspect = sizes().width / sizes().height
                camera.updateProjectionMatrix()
                renderer.setSize(sizes().width , sizes().height)
            }
            window.addEventListener('resize', resizeHandler)

            const animate = () => {
                window.requestAnimationFrame(animate)
                controls.update()

                // update the Moon's position and rotation
                renderer.render(scene, camera)
            }

            animate()
        }
    });


    return (
        <main>
            <div class="text">
                <div class="welcome-text">Hello, let's share your skills!</div>
                <button>Create CV</button>
            </div>
            <div class='scene' ref={sceneRef}/>
        </main>
    )
}

export default Welcome;