import React from 'react';
import ReactDOM from 'react-dom';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import { withStyles } from '@material-ui/core/styles';
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';
import MaterialInfo from './materialinfo.js';
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

const LINKS = 'GET_LINKS';
const INFO = 'GET_INFORMATION';

class MaterialHome extends React.Component {
    constructor(props) {
        super(props);
        this.state = {option: LINKS, url: '', info: null, submit: false};
        this.handleTextChange = this.handleTextChange.bind(this);
        this.handleSelectChange = this.handleSelectChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleSubmit(event) {
        event.preventDefault();
        if (this.state.url === '') return;
        const req = new XMLHttpRequest();
        switch (this.state.option) {
            case LINKS:
                req.open('POST', 'http://127.0.0.1:3000/links');
                req.onreadystatechange = () => {
                    if (req.readyState === 4 && req.status === 200) {
                        const response = JSON.parse(req.responseText);
                        console.log(response);
                    }
                };
                req.send(JSON.stringify(this.state));
                break;
            case INFO:
                req.open('POST', 'http://127.0.0.1:3000/info');
                req.onreadystatechange = () => {
                    if (req.readyState === 4 && req.status === 200) {
                        const response = JSON.parse(req.responseText);
                        this.setState({info: response});
                        this.setState({submit: true});
                        console.log(response);
                    }
                };
                req.send(JSON.stringify(this.state));
                break;
        }
    }
    
    handleTextChange(event) {
        this.setState({url: event.target.value});
    }

    handleSelectChange(event) {
        this.setState({option: event.target.value});
    }

    render() {
        if (!this.state.submit) {
            return (
                <form>
                    <StyledTextField label="URL" onChange={this.handleTextChange} fullWidth={true}/>
                    <br/>
                    <StyledSelect value={this.state.option} onChange={this.handleSelectChange}>
                        <MenuItem value={LINKS}>Get Links</MenuItem>
                        <MenuItem value={INFO}>Get Information</MenuItem>
                    </StyledSelect>
                    <br/>
                    <Button onClick={this.handleSubmit} variant="contained" color="primary">
                        Submit
                    </Button>
                </form>
            );
        }
        switch (this.state.option) {
            case INFO:
                return <MaterialInfo info={this.state.info}/>
            default:
                console.log("Invalid option.");
        }
    }
}

export default MaterialHome;
