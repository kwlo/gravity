import React from 'react';
import { Router } from "@reach/router";

import LandingPage from './components/LandingPage';
import Blank from './components/Blank';

export default function () {
  return (
    <Router>
      <LandingPage path="/" />
      <Blank path="blank" />
    </Router>
  );

};
