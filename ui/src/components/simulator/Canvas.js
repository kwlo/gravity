import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Engine, Scene } from '@babylonjs/core';

const useStyles = makeStyles((theme) => ({
  root: {
    height: '100%',
    overflow: 'hidden',
    width: '100%',
    padding: 0,
    margin: 0
  },
  canvas: {
    height: '100%',
    width: '100%',
    touchAction: 'none'
  }
}));

const useSetup = (canvasRef, onRender, onSetupScene) => {
  React.useEffect(() => {
    const engine = new Engine(
      canvasRef.current,
      true
    );
    const scene = new Scene(engine);

    if (onSetupScene) {
      onSetupScene(scene);
    }
    
    engine.runRenderLoop(() => {
      if (onRender) {
        onRender(scene);
      }
      scene.render();
    });
    
    const resize = () => {
      if (scene) {
          scene.getEngine().resize();
      }
    }
    window.addEventListener('resize', resize);

    return () => {
      if (scene !== null) {
        scene.dispose();
      }
      window.removeEventListener('resize', resize);
    }
  }, [canvasRef, onRender, onSetupScene]);
};

const Canvas = ({onRender, onSetupScene}) => {
  const classes = useStyles();
  const canvasRef = React.useRef();

  useSetup(canvasRef, onRender, onSetupScene);

  return (
    <div className={classes.root}>
      <canvas ref={canvasRef} className={classes.canvas} />
    </div>
  );
};

export default Canvas;
