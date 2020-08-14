import React from 'react';
import { Router } from "@reach/router";
import { makeStyles } from '@material-ui/core/styles';

import LandingPage from './components/LandingPage';
import Blank from './components/Blank';

const useStyles = makeStyles((theme) => ({
  router: {
    height: '100%',
    width: '100%'
  }
}));

export default function () {
  const classes = useStyles();

  return (
    <Router className={classes.router}>
      <LandingPage path="/" />
      <Blank path="blank" />
    </Router>
  );

};
