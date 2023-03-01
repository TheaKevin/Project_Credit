import './App.css';
import { useEffect, useState } from 'react';
import { Route, Routes, HashRouter, Link, Navigate } from 'react-router-dom';
import { FaMoneyBill, FaRegClipboard, FaUserCircle } from 'react-icons/fa';
import { Sidebar, Menu, MenuItem, SubMenu, useProSidebar, menuClasses } from 'react-pro-sidebar';
import 'bootstrap/dist/css/bootstrap.min.css'
import { ChecklistPencairan } from './pages/ChecklistPencairan';
import { Laporan } from './pages/Laporan';
import { Login } from './pages/Login';
import { ChangePasswordModal } from './pages/ChangePasswordModal';
import logo from "./assets/logoSidebar.png"
import axios from 'axios';
import { Loading } from './pages/Loading';

function App() {
  const { collapseSidebar } = useProSidebar();
  const [loading, setLoading] = useState(true);
  const [isLoggedIn, setLogin] = useState();
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [screenHeight, setScreenHeight] = useState(0);
  const [changePasswordModal, setChangePasswordModal] = useState(false);
  const openChangePasswordModal = () => setChangePasswordModal(true);
  const closeChangePasswordModal = () => setChangePasswordModal(false);

  useEffect(() => {
    (
      async () => {
        await axios.get('http://localhost:8080/user', {
          headers: {
              'Content-Type': 'application/json'
          },
          withCredentials: true
        })
        .then(res => {
          res.status === 200 ? loginSuccess(res.data.user.username, res.data.user.email) : setLogin(false);
          setLoading(false)
        })
        .catch(error => {
          setLogin(false);
          setLoading(false)
        });

        setScreenHeight(window.innerHeight);
      }
    )();
  }, [])

  const loginSuccess = (username, email) => {
    setLogin(true);
    setName(username);
    setEmail(email)
  }

  const handleLogOut = async () => {
    await axios.post('http://localhost:8080/logout', {}, {
      headers: {
          'Content-Type': 'application/json'
      },
      withCredentials: true
    })
    .then(res => {
      setLogin(false);
    })
    .catch(error => {
      setLogin(false);
    });
  }

  if(loading){
    return(
      <Loading />
    )
  }

  if (!isLoggedIn){
    return(
      <>
        <HashRouter>
          <Routes>
            <Route path="/" element={ <Login /> }/>
            <Route path="*" element={<Navigate to ="/" />}/>
          </Routes>
        </HashRouter> 
      </>
    );
  } else {
    return (
      <div style={{ display: 'flex', height: '100%', minHeight: screenHeight }}>
        <HashRouter>
          <Sidebar backgroundColor="#CC0F0F" breakPoint='lg'
          rootStyles={{
            color: 'white',
          }}>
            <Menu transitionDuration={1000} renderExpandIcon={({ open }) => <span>{open ? '-' : '+'}</span>}>
              <MenuItem className='titleSidebar' onClick={() => collapseSidebar()} icon={<img src={logo} className="logoSidebar"></img>}
              rootStyles={{
                ['.' + menuClasses.button]: {
                  '&:hover': {
                    color: '#CC0F0F',
                  },
                },
              }}>
                <b>Sinarmas</b>
              </MenuItem>
              <SubMenu className='subMenuSidebar' label="Transaksi" icon={<FaMoneyBill />}
              rootStyles={{
                ['.' + menuClasses.button]: {
                  '&:hover': {
                    color: '#CC0F0F',
                  },
                },
              }}>
                <MenuItem className='menuSidebar' routerLink={<Link to="/" />}> Checklist Pencairan </MenuItem>
              </SubMenu>
              <MenuItem className='subMenuSidebar' icon={<FaRegClipboard/>} routerLink={<Link to="/Laporan" />}
              rootStyles={{
                ['.' + menuClasses.button]: {
                  '&:hover': {
                    color: '#CC0F0F',
                  },
                },
              }}> Laporan </MenuItem>
              <SubMenu className='subMenuSidebar' label={name} icon={<FaUserCircle/>}
              rootStyles={{
                ['.' + menuClasses.button]: {
                  '&:hover': {
                    color: '#CC0F0F',
                  },
                },
              }}>
                <MenuItem className='menuSidebar' onClick={() => openChangePasswordModal()}> Change Password </MenuItem>
                <MenuItem className='menuSidebar' onClick={() => handleLogOut()}> Logout </MenuItem>
              </SubMenu>
            </Menu>
          </Sidebar>
          <main className='w-100 mainContainer'>
            <Routes>
              <Route path="/" element={ <ChecklistPencairan  /> }/>
              <Route path="/Laporan" element={ <Laporan  /> }/>
            </Routes>

            <ChangePasswordModal closeChangePasswordModal={closeChangePasswordModal} changePasswordModal={changePasswordModal} email={email} />
          </main>
        </HashRouter>
      </div>
    );
  }
}

export default App;
