import React from 'react';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles({
    pageTopContainer: {
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        marginBottom: '0.8rem',
        '@media (max-width: 768px)': {
            flexDirection: 'column',
            alignItems: 'flex-start',
        },
        '@media (max-width: 560px)': {
            padding: '0 1rem',
        },
    },
});

const PageHeader = props => {
    return <div>{props.children}</div>
}

export default PageHeader;