import type { Component } from "solid-js";
import { onMount, createSignal } from "solid-js";
import * as THREE from 'three';
import { LoadGltf } from "../Loader";
import './mainBackground.css'

const MainBackground: Component = () => {
    let sceneRef: HTMLDivElement | undefined;
    const [sizes, setSizes] = createSignal({width: 600, height: 600})
  
    onMount(async () => {
        if(sceneRef) {
            setSizes({width: window.innerWidth, height: window.innerHeight})

            // Scene
            const scene = new THREE.Scene()

            // Camera 
            const camera = new THREE.PerspectiveCamera(45, sizes().width / sizes().height, 0.1, 100)
            camera.position.y = 10
            camera.position.z = 20
            camera.lookAt(new THREE.Vector3(0, 0, 0));
            scene.add(camera)
  
            // Renderer
            const renderer = new THREE.WebGLRenderer({antialias: true})
            renderer.setSize(sizes().width, sizes().height)
            sceneRef.appendChild(renderer.domElement)
            renderer.pixelRatio = 2
            
            // Earth
            const earthMesh = await LoadGltf('Earth');
            earthMesh.scale.set(0.5, 0.5, 0.5)
            earthMesh.position.y = -1
            
            // Moon
            const moonMesh = await LoadGltf('Moon');
            moonMesh.scale.set(0.04, 0.04, 0.04)
            moonMesh.position.x = 20
            let moonAngle = 0

            // const earthMoonGroup = new THREE.Group()
            // earthMoonGroup.add(earthMesh) 
            // earthMoonGroup.add(moonMesh) 
            earthMesh.add(moonMesh)
            scene.add(earthMesh)



            // Point light
            const pointLight = new THREE.PointLight('#e8fafc', 1, 100)
            pointLight.position.set(6, 10, 10)
            scene.add(pointLight)

            // Ambient light
            const light = new THREE.AmbientLight( 0x404040 ); // soft white light
            scene.add( light );
             
            renderer.render(scene, camera)
  
            // Resize handler event listener
            const resizeHandler = (event: Event) => {
                // Update sizes
                setSizes({width: window.innerWidth, height: window.innerHeight})
  
                // Update camera aspect and renderer size 
                camera.aspect = sizes().width / sizes().height
                camera.updateProjectionMatrix()
                renderer.setSize(sizes().width , sizes().height)
            }
            window.addEventListener('resize', resizeHandler)
  
            const animate = () => {
                window.requestAnimationFrame(animate)

                // rotate the Earth
                earthMesh.rotation.y += 0.01;

                // update the Moon's position and rotation
                moonAngle += 0.02;
                moonMesh.position.set(
                    20 * Math.cos(moonAngle),
                    0,
                    20 * Math.sin(moonAngle)
                );
                moonMesh.rotation.y += 0.02;
                renderer.render(scene, camera)
            }
  
            animate()
        }
    });
  
  //   onCleanup(() => {
  //     window.removeEventListener('resize', resizeHandler)
  //   })
  
  
    return (
      <div class='scene' ref={sceneRef}/>
    );
  }

export default MainBackground