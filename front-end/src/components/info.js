import React from 'react';

class Info extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            info: props.info,
            websocket: props.websocket
        };
    }

    render() {
        return (
            <React.Fragment>
                <h1>INFO</h1>
                <div>{this.state.info}</div>
            </React.Fragment>
        );
    }
}

export default Info;
