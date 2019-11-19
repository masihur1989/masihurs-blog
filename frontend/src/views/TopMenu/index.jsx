import React from 'react';
import { HashRouter, Route, Switch } from 'react-router-dom';
// import { UserInfo } from '../../store/user/reducers';
import i18n from '../../i18n/i18n';

import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';

import HomeComponent from '../homepage/index';
import DashboardComponent from '../dashboard/index';
import LandingPage from '../LandingPage/LandingPage';

const useStyles = makeStyles({
    root: {
        flexGrow: 1,
        backgroundColor: '#fff',
        borderBottom: '1px solid #CACFD8',
    },
    menuIconContainer: {
        width: '5rem',
        textAlign: 'center',
        '@media (max-width: 540px)': {
            width: 'initial',
        },
    },
    menuIcon: {
        fill: '#CACFD8',
    },
    navItem: {
        height: '100%',
    },
    logo: {
        flexGrow: 1,
        lineHeight: 1,
        '@media (max-width: 540px)': {
            margin: '0',
        },
        '& a': { display: 'inline-block' },
        '& svg': {
            verticalAlign: 'middle',
            marginRight: '.4rem',
        },
        '& .site-name': {
            verticalAlign: 'middle',
            color: '#f74527',
            fontSize: '1.1875rem',
            lineHeight: '1.1875rem',
        },
    },
    toolbar: {
        minHeight: '3.5625rem',
        '@media (min-width: 600px)': {
            height: '4rem',
        },
        '@media (min-width: 0px) and (orientation: landscape)': {
            height: '3rem',
        },
    },
    support: {
        '@media (max-width: 740px)': {
            display: 'none',
        },
    },
});


class TopMenu extends React.Component {

    render() {
        return (
            <Box>
                <HashRouter>
                    <Switch>
                        <Route
                            path="/"
                            component={HomeComponent}
                            exact
                        />
                        <Route
                            path="/dashboard"
                            component={DashboardComponent}
                        />
                        <Route 
                            path="/landing-page" 
                            component={LandingPage} 
                        />
                    </Switch>
                </HashRouter>
            </Box>
        );
    }
}

export default TopMenu;