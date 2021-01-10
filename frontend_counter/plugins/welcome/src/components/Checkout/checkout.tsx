import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import {Cookies} from '../../Cookie'
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntPromotion } from '../../api/models/EntPromotion'; // import interface User
import { EntTypeRoom } from '../../api/models/EntTypeRoom'; // import interface Resolution
import { EntStatusRoom } from '../../api/models/EntStatusRoom'; // import interface Playlist
// me
import { EntCheckIn } from '../../api/models/EntCheckIn'; // import interface checkin
import { EntStatus } from '../../api/models/EntStatus'; // import interface Status
import { EntCounterStaff } from '../../api/models/EntCounterStaff'; // import interface CounterStaff
// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

/*
interface DataRoom {
  Promotion: number;
  StatusRoom: number;
  TypeRoom: number;
  Price: number;
  RoomNumber: string;
  // create_by: number;
}
*/
interface CheckOut {
    CheckinsID: number;
    StatussID: number;
    CounterstaffsID: number;
}

const checkout: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();

  var ck = new Cookies()
  var cookieName = ck.GetCookie()

  const [CheckOut, setDataRoom] = React.useState<Partial<CheckOut>>({});
  const [StatusRoom, setStatusRoom] = React.useState<EntStatusRoom[]>([]);
  const [Promotion, setPromotion] = React.useState<EntPromotion[]>([]);
  const [TypeRoom, setTypeRoom] = React.useState<EntTypeRoom[]>([]);

  // checkout
  const [Checkin, setCheckin] = React.useState<EntCheckIn[]>([]);  
  const [Status, setStatus] =   React.useState<EntStatus[]>([]);
  const [CounterStaff, setCounterStaff] =   React.useState<EntCounterStaff[]>([]);
  //

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });


  // tata

  const getTypeRoom = async () => {
    const res = await api.listTyperoom({ limit: 10, offset: 0 });
    setTypeRoom(res);
  };
  const getPromotion = async () => {
    const res = await api.listPromotion({ limit: 10, offset: 0 });
    setPromotion(res);
  };

  // set data for checkout
  const getcheckIn = async () => {
    const res = await api.listCheckin({ limit: 10, offset: 0 });
    setCheckin(res);
  };

  const getstatus = async () => {
    const res = await api.listStatus({ limit: 10, offset: 0 });
    setStatus(res);
  };
  
  
  const getstaff = async () => {
    const res = await api.listCounterStaff({ limit: 10, offset: 0 });
    setCounterStaff(res);
  };
  

  // Lifecycle Hooks
  useEffect(() => {

    getPromotion();
    getTypeRoom();

    // me 
    getcheckIn();
    getstatus();
    getstaff();
  }, []);

  // set data to object DataRoom
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof CheckOut;
    const { value } = event.target;
    setDataRoom({ ...CheckOut, [name]: value });
    console.log(CheckOut);
  };
 
  
  
  // clear input form
  function clear() {
    setDataRoom({});
  }

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/checkouts';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(CheckOut),
    };

    console.log(CheckOut); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.id != null) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }

  function Clears() {
    ck.ClearCookie()
    window.location.reload(false)
  }

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`Watch Video`}>
        <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
        <div style={{ marginLeft: 10, marginRight:20 }}>{cookieName}</div>
        <Button
          variant="outlined"
          color="secondary"
          size="large"
          onClick={Clears}
          >
          Logout
        </Button>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              
            </Grid>
            <Grid item xs={9}>
            
           </Grid>
            
            <Grid item xs={3}>
              <div className={classes.paper}>check in</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือก</InputLabel>
                <Select
                  name="CheckinsID"
                  value={CheckOut.CheckinsID  || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {Checkin.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.id}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>status</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือก</InputLabel>
                <Select
                  name="StatussID"
                  value={CheckOut.StatussID  || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {Status.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.description}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>counter staff</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือก</InputLabel>
                <Select
                   name="CounterstaffsID"
                   value={CheckOut.CounterstaffsID || ''} // (undefined || '') = ''
                   onChange={handleChange}
                >
                  {CounterStaff.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              
            </Grid>
            <Grid item xs={9}>
           
           </Grid>
           <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                check out
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default checkout;