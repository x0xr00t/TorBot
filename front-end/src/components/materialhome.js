import React from 'react';
import ReactDOM from 'react-dom';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import { withStyles } from '@material-ui/core/styles';

const StyledTextField = withStyles({
    root: {
        'background-color': 'white',
        margin: 15
    }
})(TextField);

class MaterialHome extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <form>
                <StyledTextField/>
                <br/>
                <StyledTextField/>
                <br/>
                <StyledTextField/>
                <br/>
                <StyledTextField/>
                <br/>
                <Button variant="contained" color="primary">
                    Hello World
                </Button>
            </form>
        );
    }
}

export default MaterialHome;
