import React from 'react';
import ReactDOM from 'react-dom';
import 'antd/dist/antd.css';
import './index.css';
import * as serviceWorker from './serviceWorker';

import { Player } from '../src/pages/Player/PlayerPage';
import { Header } from '../src/components/Header/Header';
import { ApolloProvider } from '@apollo/client';
import { client } from './util/graph';
import { AuthAppRouter } from './App';
import { Switch, Route, BrowserRouter as Router } from 'react-router-dom';
import { DemoContext } from './DemoContext';

ReactDOM.render(
    <React.StrictMode>
        <ApolloProvider client={client}>
            <Router>
                <Switch>
                    <Route path="/demo" exact>
                        <DemoContext.Provider value={{ demo: true }}>
                            <Header />
                            <Player />
                        </DemoContext.Provider>
                    </Route>
                    <Route path="/">
                        <DemoContext.Provider value={{ demo: false }}>
                            <AuthAppRouter />
                        </DemoContext.Provider>
                    </Route>
                </Switch>
            </Router>
        </ApolloProvider>
    </React.StrictMode>,
    document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
