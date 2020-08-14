import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import * as BABYLON from '@babylonjs/core';
import Canvas from './Canvas';

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100%',
    height: '100%'
  },
  info: {
    position: 'absolute',
    zIndex: 1000,
    color: 'red'
  }
}));

let angle = 0;
const handleRender = (infoRef) => (engine, scene) => {
  const cot = scene.getNodeByID("light2Node");
  
  scene.registerBeforeRender(() => {
    cot.rotation.y = angle;
    angle = (angle + 0.0001) % 4;
  });

  infoRef.current.innerHTML = `FPS: ${engine.getFps().toFixed()}`;
};

const handleSetupScene = (scene) => {
  const camera = new BABYLON.FreeCamera("camera1", new BABYLON.Vector3(0, 5, -10), scene);
  camera.setTarget(BABYLON.Vector3.Zero());
  
  // Add lights
  new BABYLON.HemisphericLight(
    "light1",
    new BABYLON.Vector3(1, 1, 0),
    scene
  );
  
  const light = new BABYLON.PointLight(
    "light2",
    new BABYLON.Vector3(0, 1, -1),
    scene
  );

  const cot = new BABYLON.TransformNode("light2Node"); 

  light.parent = cot;

  // Add sphere
  BABYLON.MeshBuilder.CreateSphere(
    "sphere",
    { diameter: 2 },
    scene
  );
};

const Simulator = () => {
  const classes = useStyles();
  const infoRef = React.useRef(null);

  return (
    <div className={classes.root}>
      <div ref={infoRef} className={classes.info} />
      <Canvas
        onRender={handleRender(infoRef)}
        onSetupScene={handleSetupScene}
        onDestroy={() => console.log('destroyed')}
      />
    </div>
  );
};

export default Simulator;
