import PropTypes from 'prop-types';
import { useEffect, useState } from 'react';

import React from 'react';
import {
  Card,
  Box,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TableContainer
} from '@mui/material';
import { Scrollbar } from 'src/components/scrollbar';

export const CheckpointState = (props) => {
  const { items = [] } = props;
  const [selectedEqNo, setSelectedEqNo] = useState(null);

  const sortedItems = items.sort((a, b) => b.id - a.id);
  const uniqueItemsMap = new Map();
  sortedItems.forEach(item => {
    if (!uniqueItemsMap.has(item.inspeqno)) {
      uniqueItemsMap.set(item.inspeqno, item);
    }
  });

  const uniqueItems = Array.from(uniqueItemsMap.values());

  const handleRowClick = (eqNo) => {
    if (selectedEqNo === eqNo) {
      setSelectedEqNo(null);
    } else {
      setSelectedEqNo(eqNo);
    }
  };

  const selectedItemDetails = items.filter(item => item.inspeqno === selectedEqNo && item.id !== uniqueItemsMap.get(selectedEqNo).id);


  const getStatusText = (status) => {
    switch (status) {
      case "0":
        return '정상 대기';
      case "1":
        return '검색시작';
      case "2":
        return '검색중';
      case "3":
        return '분석중';
      case "4":
        return '검색종료';
      case "5":
        return '이상오류';
      default:
        return '알 수 없는 상태';
    }
  };

  const getStatusColor = (status) => {
    switch (status) {
      case "0":
        return '#28a745';
      case "1":
        return '#17a2b8';
      case "2":
        return '#007bff';
      case "3":
        return '#5816db';
      case "4":
        return '#ffc107';
      case "5":
        return '#dc3545';
      default:
        return '#6c757d';
    }
  };
  return (
    <Card>
      <Scrollbar>
        <Box sx={{ minWidth: 800 }}>
          <TableContainer sx={{ maxHeight: 420 }}>
            <Table stickyHeader>
              <TableHead>
                <TableRow>
                  <TableCell>Equipment No</TableCell>
                  <TableCell>Equipment Status</TableCell>
                  <TableCell>Date</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {uniqueItems.map((item) => (
                  <React.Fragment key={item.id}>
                    <TableRow onClick={() => handleRowClick(item.inspeqno)}>
                      <TableCell>{item.inspeqno}</TableCell>
                      <TableCell>
                        <span
                          style={{
                            backgroundColor: getStatusColor(item.inspeqstatus),
                            padding: '0.2rem 0.5rem',
                            borderRadius: '0.2rem',
                            fontWeight: 'bold',
                            color: 'white'
                          }}
                        >
                          {getStatusText(item.inspeqstatus)}
                        </span>
                      </TableCell>
                      <TableCell>{item.qdate}</TableCell>
                    </TableRow>
                    {selectedEqNo === item.inspeqno && (
                      <TableRow>
                        <TableCell colSpan={3}>
                          <Table>
                            <TableHead>
                              <TableRow>
                                <TableCell>Equipment No</TableCell>
                                <TableCell>Equipment Status</TableCell>
                                <TableCell>Date</TableCell>
                              </TableRow>
                            </TableHead>
                            <TableBody>
                              {selectedItemDetails.map(detail => (
                                <TableRow key={detail.id}>
                                  <TableCell>{detail.inspeqno}</TableCell>
                                  <TableCell>
                                    <span
                                      style={{
                                        backgroundColor: getStatusColor(detail.inspeqstatus),
                                        padding: '0.2rem 0.5rem',
                                        borderRadius: '0.2rem',
                                        fontWeight: 'bold',
                                        color: 'white'
                                      }}
                                    >
                                      {getStatusText(detail.inspeqstatus)}
                                    </span>
                                  </TableCell>
                                  <TableCell>{detail.qdate}</TableCell>
                                </TableRow>
                              ))}
                            </TableBody>
                          </Table>
                        </TableCell>
                      </TableRow>
                    )}
                  </React.Fragment>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Box>
      </Scrollbar>
    </Card>
  );
};

CheckpointState.propTypes = {
  items: PropTypes.arrayOf(PropTypes.shape({
    id: PropTypes.number,
    inspEqNo: PropTypes.string,
    inspEqStatus: PropTypes.string,
  })),
};
