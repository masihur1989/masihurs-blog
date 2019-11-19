import { createMuiTheme } from '@material-ui/core/styles'

const theme = createMuiTheme({
    typography: {
        lineHeight: 1,
        fontFamily: 'nunitoregular, sans-serif',
        htmlFontSize: 18,
        fontStyle: 'normal',
        fontWeight: '100',
        useNextVariants: true,
        body1: {
            fontFamily: 'nunitosemibold, sans-serif',
            fontSize: 16,
            lineHeight: 1.4,
            '@media (max-width: 1050px)': {
                fontSize: '0.9375rem',
            },
        },
        body2: {
            fontSize: 14,
            lineHeight: 1.2,
            textTransform: 'initial',
        },
        h1: {
            fontSize: 21,
            lineHeight: 1.31,
        },
        h2: {
            fontFamily: 'nunitosemibold, sans-serif',
            fontSize: 20,
            lineHeight: 1.31,
        },
        h3: {
            fontSize: 18,
            lineHeight: 1.31,
        },
        h5: {
            fontSize: 14,
            lineHeight: 1.31,
        },
        h6: {
            fontSize: 12,
            lineHeight: 1.31,
        },
    },
    shadows: Array(25).fill('none'),
    spacing: value => value ** 2,
    palette: {
        primary: {
            main: '#1ba4f1',
            contrastText: '#0b143b',
        },
        secondary: {
            main: '#c0c4ca',
            contrastText: '#39397C',
        },
        textPrimary: {
            main: '#000000',
            fontWeight: '600',
        },
        grey: {
            main: 'red',
        },
    },
    // overrides: {
    //   MuiTableSortLabel: {
    //     iconDirectionAsc: Alarm,
    //   },
    // },
})

export default theme;