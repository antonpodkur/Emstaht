import { GLTFLoader } from 'three/addons/loaders/GLTFLoader.js';

export async function LoadGltf(modelName: string): Promise<THREE.Group> {
    const loader = new GLTFLoader();
    const gltf = await loader.loadAsync('/src/assets/'+modelName+'.glb');
    return gltf.scene;
}