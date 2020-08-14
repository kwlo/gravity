import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import * as BABYLON from '@babylonjs/core';
import Canvas from './Canvas';

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100%',
    height: '100%'
  }
}));

const handleRender = (scene) => {
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
  
  new BABYLON.PointLight(
    "light2",
    new BABYLON.Vector3(0, 1, -1),
    scene
  );

  // Add sphere
  BABYLON.MeshBuilder.CreateSphere(
    "sphere",
    { diameter: 2 },
    scene
  );
};

const Simulator = () => {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Canvas
        onRender={handleRender}
        onSetupScene={handleSetupScene}
      />
    </div>
  );
};

export default Simulator;
