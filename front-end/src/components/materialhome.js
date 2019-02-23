import React from 'react';
import ReactDOM from 'react-dom';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import { withStyles } from '@material-ui/core/styles';
import './materialhome.css';

const StyledTextField = withStyles({
    root: {
        'background-color': 'white',
        'margin': 15,
        'padding': 5,
        'border-radius': 6
    }
})(TextField);

class MaterialHome extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <form>
                <StyledTextField label="URL"/>
                <br/>
                <Button type="submit" variant="contained" color="primary">
                    Submit
                </Button>
            </form>
        );
    }
}

export default MaterialHome;
