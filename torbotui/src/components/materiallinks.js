import React from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableRow from '@material-ui/core/TableRow';
import TableHead from '@material-ui/core/TableHead';
import TableCell from '@material-ui/core/TableCell';
import Paper from '@material-ui/core/Paper';
import { Tab } from '@material-ui/core';

let ws;

let id = 0;
function createRow(link, status) {
    id += 1;
    return {id, link, status};
}
class MaterialLinks extends React.Component {
    constructor(props) {
        super(props);
        this.state = {linkStatus: []};
        this.state.linkStatus.push(createRow('http://www.google.com', 'GOOD'));
        ws = new WebSocket('ws://127.0.0.1/links?url=' + encodeURIComponent(props.url));
        ws.onmessage = this.handleMessage.bind(this); 
    }

    handleMessage(msg) {
        
    }

    render() {
        return (
            <Paper>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>Link</TableCell>
                            <TableCell>Status</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {this.state.linkStatus.map(linkStatus => (
                            <TableRow key={linkStatus.id}>
                                <TableCell>{linkStatus.link}</TableCell>
                                <TableCell>{linkStatus.status}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </Paper>
        );
    }
}

export default MaterialLinks;