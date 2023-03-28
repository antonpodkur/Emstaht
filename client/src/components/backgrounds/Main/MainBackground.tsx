import type { Component, Ref } from "solid-js";
import { onMount, createSignal } from "solid-js";
import * as THREE from 'three';
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls'
import './mainBackground.css'

const MainBackground: Component = () => {
    let sceneRef: undefined; 
    const [sizes, setSizes] = createSignal({width: 600, height: 600})
  
    onMount(() => {
      if(sceneRef) {
          setSizes({width: window.innerWidth, height: window.innerHeight})
          
          // Scene
          const scene = new THREE.Scene()
          const geometry = new THREE.SphereGeometry(4, 64, 64)
          const material = new THREE.MeshStandardMaterial({
              color: "#8d42f5"
          })
          const mesh = new THREE.Mesh(geometry, material)
          scene.add(mesh)
          
  
          // Light
          const light = new THREE.PointLight('white', 1, 100)
          light.position.set(0, 10, 10)
          scene.add(light)
          
          // Camera 
          const camera = new THREE.PerspectiveCamera(45, sizes().width / sizes().height, 0.1, 100)
          camera.position.z = 20
          scene.add(camera)
  
          // Renderer
          const renderer = new THREE.WebGLRenderer({antialias: true})
          renderer.setSize(sizes().width, sizes().height)
          sceneRef.appendChild(renderer.domElement)
          renderer.pixelRatio = 2
          renderer.render(scene, camera)
  
          // Controls
          const controls = new OrbitControls(camera, renderer.domElement)
          controls.enableDamping = true
          controls.enablePan = false
          controls.enableZoom = false
          controls.autoRotate = true
          controls.autoRotateSpeed = 5
  
          // Resize handler event listener
          const resizeHandler = (event: Event) => {
              // Update sized
              setSizes({width: window.innerWidth, height: window.innerHeight})
  
              // Update camera aspect and renderer size 
              camera.aspect = sizes().width / sizes().height
              camera.updateProjectionMatrix()
              renderer.setSize(sizes().width , sizes().height)
          }
          window.addEventListener('resize', resizeHandler)
  
          const animate = () => {
              controls.update()
              renderer.render(scene, camera)
              window.requestAnimationFrame(animate)
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