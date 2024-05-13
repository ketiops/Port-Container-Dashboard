import PropTypes from 'prop-types';
import React from 'react';
import { useEffect, useState } from 'react';
import {
  Card,
  Box,
  TextField,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TableContainer, Button, Dialog, DialogTitle, DialogContent, DialogActions
} from '@mui/material';
import { Scrollbar } from 'src/components/scrollbar';
import axiosInstance from "src/axiosInstance";


export const ContainerSpec = (props) => {
  const [items, setItems] = useState([]);
  const [searchValue, setSearchValue] = useState("");
  const [remarkValue, setRemarkValue] = useState("");

  const [selectedRow, setSelectedRow] = useState(null);
  const [remarkData, setRemarkData] = useState([]);

  const handleRowClick = (itemId) => {
    if (selectedRow === itemId) {
      setSelectedRow(null);
    } else {
      setSelectedRow(itemId);
      axiosInstance.get("/remarks?informationid=eq." + itemId.toString())
        .then((response) => setRemarkData(response.data))
        .catch((error) => console.error('Error fetching data: ', error));
    }
  }

  const [open, setOpen] = useState(false);
  const handleOpen = () => {
    setOpen(true);
  };
  const handleClose = () => {
    setOpen(false);
  };
  const handleCloseWithAdd = () => {
    const postData = {
      "inspremark": remarkValue,
      "informationid": selectedRow
    };
    axiosInstance.post("/remarks", postData)
      .then(response => {
        setOpen(false);
      })
      .catch(error => {
        console.error("Error inserting data:", error);
      });
    setRemarkValue("")
  };

  const [sortOrder, setSortOrder] = useState("default");

  const handleSortClick = () => {
    if (sortOrder === "default") {
      setSortOrder("desc");
    } else if (sortOrder === "desc") {
      setSortOrder("asc");
    } else {
      setSortOrder("default");
    }
  };

  let sortedItems = [...items];
  if (sortOrder === "asc") {
    sortedItems.sort((a, b) => new Date(a.qdate) - new Date(b.qdate));
  } else if (sortOrder === "desc") {
    sortedItems.sort((a, b) => new Date(b.qdate) - new Date(a.qdate));
  }

  useEffect(() => {
    let url = "/containerspec";
    if (searchValue.trim()) {
      url += `?inspno=like.${searchValue}*`;
    }
    axiosInstance.get(url)
      .then((response) => setItems(response.data))
      .catch((error) => console.error('Error fetching data: ', error));
  }, [searchValue]);

  return (
    <Card>
      <Box sx={{ padding: 2 }}>
        <TextField
          label="Search by Insp No"
          variant="outlined"
          fullWidth
          value={searchValue}
          onChange={(e) => setSearchValue(e.target.value)}
        />
      </Box>
      <Scrollbar>
        <Box sx={{ minWidth: 800 }}>
          <TableContainer sx={{ maxHeight: 420 }}>
            <Table stickyHeader>
              <TableHead>
                <TableRow>
                  <TableCell>
                    Equipment No
                  </TableCell>
                  <TableCell>
                    Inspection No
                  </TableCell>
                  <TableCell>
                    Start Time
                  </TableCell>
                  <TableCell>
                    End Time
                  </TableCell>
                  <TableCell>
                    Package Match
                  </TableCell>
                  <TableCell>
                    Inspection Result CD
                  </TableCell>
                  <TableCell>
                    Detection Count
                  </TableCell>
                  <TableCell>
                    Fault CD
                  </TableCell>
                  <TableCell>
                    Result Image Dir
                  </TableCell>
                  <TableCell onClick={handleSortClick}>Date</TableCell>

                </TableRow>
              </TableHead>
              <TableBody>
                {sortedItems.map((item) => (
                  <React.Fragment key={item.id}>
                    <TableRow key={item.id}
                      onClick={() => handleRowClick(item.id)}
                      sx={{
                        backgroundColor: selectedRow === item.id ? 'rgba(0, 0, 255, 0.1)' : 'inherit',
                        fontWeight: selectedRow === item.id ? 'bold' : 'normal'
                      }}>
                      <TableCell>{item.inspeqno}</TableCell>
                      <TableCell>{item.inspno}</TableCell>
                      <TableCell>{item.inspstarttime}</TableCell>
                      <TableCell>{item.inspendtime}</TableCell>
                      <TableCell>{item.pckmatch}</TableCell>
                      <TableCell>{item.insprsltcd}</TableCell>
                      <TableCell>{item.detectioncnt}</TableCell>
                      <TableCell>{item.faultcd}</TableCell>
                      <TableCell>{item.insprsltimgdir}</TableCell>
                      <TableCell>{item.qdate}</TableCell>
                    </TableRow >
                    {selectedRow === item.id && (
                      <TableRow>
                        <TableCell colSpan={9}>
                          <Table>
                            <TableHead>
                              <TableRow>
                                <TableCell>Insp Remark</TableCell>
                                <TableCell align="right">
                                  <Button variant="contained" color="primary" onClick={handleOpen}>
                                    +Add remark
                                  </Button>
                                  <Dialog open={open} onClose={handleClose}>
                                    <DialogTitle>Add Remark</DialogTitle>
                                    <DialogContent>
                                      <TextField value={remarkValue} label="Enter the Remark"
                                        variant="outlined"
                                        fullWidth
                                        onChange={(e) => setRemarkValue(e.target.value)}
                                      ></TextField>
                                    </DialogContent>
                                    <DialogActions>
                                      <Button onClick={handleClose} color="primary">
                                        Cancel
                                      </Button>
                                      <Button onClick={handleCloseWithAdd} color="primary">
                                        Add
                                      </Button>
                                    </DialogActions>
                                  </Dialog>
                                </TableCell>
                              </TableRow>
                            </TableHead>
                            <TableBody>
                              {remarkData.map((remark) => (
                                <TableRow key={remark.id}>
                                  <TableCell>{remark.inspremark}</TableCell>
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
    </Card >
  );
};