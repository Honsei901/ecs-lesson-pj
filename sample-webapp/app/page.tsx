'use client';

import React, { Fragment, useState } from 'react';
import { HealthCheckButton } from './ui/HealthCheckButton';
import { ApiButton } from './ui/ApiButton';

const Home = () => {
  const [text, setText] = useState('Hello, world!');
  return (
    <Fragment>
      <div className="flex flex-col">
        <div className="flex items-center gap-12">
          <HealthCheckButton setText={setText} />
          <ApiButton setText={setText} />
        </div>
        <div className="mt-20 flex justify-center">
          <h2 className="text-3xl">{text}</h2>
        </div>
      </div>
    </Fragment>
  );
};

export default Home;
