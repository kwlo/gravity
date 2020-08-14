import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Engine, Scene } from '@babylonjs/core';

const useStyles = makeStyles(() => ({
  root: {
    height: '100%',
    overflow: 'hidden',
    width: '100%',
    padding: 0,
    margin: 0,
  },
  canvas: {
    height: '100%',
    width: '100%',
    touchAction: 'none',
  },
}));

const useSetup = (canvasRef, onRender, onSetupScene, onDestroy) => {
  React.useEffect(() => {
    const engine = new Engine(canvasRef.current, true);
    const scene = new Scene(engine);

    if (onSetupScene) {
      onSetupScene(scene);
    }

    engine.runRenderLoop(() => {
      if (onRender) {
        onRender(engine, scene);
      }
      scene.render();
    });

    const resize = () => {
      if (scene) {
        scene.getEngine().resize();
      }
    };
    window.addEventListener('resize', resize);

    return () => {
      window.removeEventListener('resize', resize);

      if (onDestroy) {
        onDestroy(engine, scene);
      }

      if (scene) {
        scene.dispose();
      }
      if (engine) {
        engine.dispose();
      }
    };
  }, [canvasRef, onRender, onSetupScene, onDestroy]);
};

const Canvas = ({ onRender, onSetupScene, onDestroy }) => {
  const classes = useStyles();
  const canvasRef = React.useRef();

  useSetup(canvasRef, onRender, onSetupScene, onDestroy);

  return (
    <div className={classes.root}>
      <canvas ref={canvasRef} className={classes.canvas} />
    </div>
  );
};

export default React.memo(Canvas);
