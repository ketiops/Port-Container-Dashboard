import PropTypes from 'prop-types';
import {
  Card,
  Box,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Checkbox,
  TableContainer
} from '@mui/material';
import { Scrollbar } from 'src/components/scrollbar';

export const CheckpointInfo = (props) => {
  const { items = [] } = props;

  return (
    <Card>
      <Scrollbar>
        <Box sx={{ minWidth: 800 }}>
          <TableContainer sx={{ maxHeight: 420 }}>
            <Table stickyHeader>
              <TableHead>
                <TableRow>
                  <TableCell>Equipment No</TableCell>
                  <TableCell>Inspection Auto</TableCell>
                  <TableCell>Inspection Name</TableCell>
                  <TableCell>Inspection Location</TableCell>
                  <TableCell>Contact</TableCell>
                  <TableCell>Date</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {items.map((item) => (
                  <TableRow key={item.id}>
                    <TableCell>{item.inspeqno}</TableCell>
                    <TableCell>{item.inspauto}</TableCell>
                    <TableCell>{item.inspname}</TableCell>
                    <TableCell>{item.insploc}</TableCell>
                    <TableCell>{item.inspcontact}</TableCell>
                    <TableCell>{item.qdate}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Box>
      </Scrollbar>
    </Card>
  );
};

CheckpointInfo.propTypes = {
  items: PropTypes.arrayOf(PropTypes.shape({
    id: PropTypes.number,
    inspEqNo: PropTypes.string,
    inspAuto: PropTypes.string,
    inspName: PropTypes.string,
    inspLoc: PropTypes.string,
    inspContact: PropTypes.string,
  })),
};