import {
  Card,
  Box,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TableContainer,
  Select,
  MenuItem,
  InputLabel,
  FormControl,
  TextField,
  Grid,
  FormControlLabel,
  Checkbox
} from '@mui/material';
import { useEffect, useState } from 'react';
import axiosInstance from "src/axiosInstance";

import { Scrollbar } from 'src/components/scrollbar';

export const PreInformation = (props) => {
  const { items: defaultItems = [] } = props;

  const [data, setData] = useState(defaultItems);
  const [searchType, setSearchType] = useState('all');
  const [searchText, setSearchText] = useState('');
  const [filteredItems, setFilteredItems] = useState(defaultItems);
  const [checkBoxes, setCheckBoxes] = useState({
    localData: false,
    COPINO: false,
    COARRI: false
  });
  useEffect(() => {
    axiosInstance.get(searchType === 'all' ? "/preinformation" : `/preinformation?${searchType}=ilike.${searchText}*`)
      .then(response => {
        setData(response.data);
      })
      .catch(error => console.error('Error fetching data: ', error));
  }, [searchType, searchText]);

  useEffect(() => {
    let filteredData = data;

    if (searchType !== 'all') {
      filteredData = filteredData.filter(item => item[searchType] && item[searchType].includes(searchText));
    }

    if (checkBoxes.localData || checkBoxes.COPINO || checkBoxes.COARRI) {
      filteredData = filteredData.filter(item =>
        (checkBoxes.localData && item.typeno === '') ||
        (checkBoxes.COPINO && item.typeno === 'COPINO') ||
        (checkBoxes.COARRI && item.typeno === 'COARRI')
      );
    }

    setFilteredItems(filteredData);
  }, [data, searchType, searchText, checkBoxes]);

  return (
    <Card>
      <Box sx={{ p: 2 }}>
        <Grid container alignItems="center" spacing={2}>
          <Grid item xs style={{ flexGrow: 1 }}>
          </Grid>
          <Grid item>
            <FormControlLabel
              control={<Checkbox checked={checkBoxes.localData} onChange={(e) => setCheckBoxes({ ...checkBoxes, localData: e.target.checked })} />}
              label="local Data"
            />
            <FormControlLabel
              control={<Checkbox checked={checkBoxes.COPINO} onChange={(e) => setCheckBoxes({ ...checkBoxes, COPINO: e.target.checked })} />}
              label="COPINO"
            />
            <FormControlLabel
              control={<Checkbox checked={checkBoxes.COARRI} onChange={(e) => setCheckBoxes({ ...checkBoxes, COARRI: e.target.checked })} />}
              label="COARRI"
            />
          </Grid>
          <Grid item xs={2}>
            <FormControl variant="outlined" fullWidth>
              <InputLabel>검색 유형</InputLabel>
              <Select
                value={searchType}
                onChange={(e) => setSearchType(e.target.value)}
                label="검색 유형"
              >
                <MenuItem value="all">전체보기</MenuItem>
                <MenuItem value="inspeqno">EQUIPMENT NO</MenuItem>
                <MenuItem value="cntrno">CONTAINER NO</MenuItem>
                <MenuItem value="truckno">TRUCK NO</MenuItem>
              </Select>
            </FormControl>
          </Grid>
          {searchType !== 'all' && (
            <Grid item xs={2}>
              <TextField
                variant="outlined"
                placeholder={`Search by ${searchType}`}
                value={searchText}
                onChange={(e) => setSearchText(e.target.value)}
                fullWidth
              />
            </Grid>
          )}
        </Grid>
      </Box>
      <Scrollbar>
        <Box sx={{ minWidth: 800 }}>
          <TableContainer sx={{ maxHeight: 420 }}>
            <Table stickyHeader>
              <TableHead>
                <TableRow>
                  <TableCell>Equipment No</TableCell>
                  <TableCell>Container No</TableCell>
                  <TableCell>Truck No</TableCell>
                  <TableCell>Type No</TableCell>
                  <TableCell>Date</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {filteredItems.map((item) => (
                  <TableRow key={item.id}>
                    <TableCell>{item.inspeqno}</TableCell>
                    <TableCell>{item.cntrno}</TableCell>
                    <TableCell>{item.truckno}</TableCell>
                    <TableCell>{item.typeno ? item.typeno : "local DB"}</TableCell>
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

