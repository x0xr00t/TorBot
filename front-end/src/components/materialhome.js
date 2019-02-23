import React from 'react';
import ReactDOM from 'react-dom';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import { withStyles } from '@material-ui/core/styles';
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';
import './materialhome.css';

const StyledTextField = withStyles({
    root: {
        'background-color': 'white',
        'margin-bottom': 5,
        'padding': 5,
        'border-radius': 6
    }
})(TextField);

const StyledSelect = withStyles({
    root: {
        'background-color': 'white',
        'margin-bottom': 5, 
        'padding': 5,
        'border-radius': 6
    }
})(Select);

class MaterialHome extends React.Component {
    constructor(props) {
        super(props);
        this.state = {option: 'GET_LINKS'};
        this.handleChange = this.handleChange.bind(this);
    }

    handleChange(event) {
        this.setState({option: event.target.value});
    }

    render() {
        return (
            <form>
                <StyledTextField label="URL"/>
                <br/>
                <StyledSelect value={this.state.option} onChange={this.handleChange}>
                    <MenuItem value={'GET_LINKS'}>Get Links</MenuItem>
                    <MenuItem value={'GET_INFORMATION'}>Get Information</MenuItem>
                </StyledSelect>
                <br/>
                <Button type="submit" variant="contained" color="primary">
                    Submit
                </Button>
            </form>
        );
    }
}

export default MaterialHome;
