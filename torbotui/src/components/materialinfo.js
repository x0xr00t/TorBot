import React from 'react';
import ReactDOM from 'react-dom';
import Table from '@material-ui/core/Table';
import TableRow from '@material-ui/core/TableRow';
import TableCell from '@material-ui/core/TableCell';

class MaterialInfo extends React.Component {
    constructor(props) {
        super(props);
        this.state = {info: props.info};
    }

    render() {
        return (
            <Table>
                <TableRow>
                    <TableCell>Something</TableCell>
                </TableRow>
            </Table>
        );
    }
}

export default MaterialInfo;