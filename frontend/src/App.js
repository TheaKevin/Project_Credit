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

function App() {
  const { collapseSidebar } = useProSidebar();
  const [isLoggedIn, setLogin] = useState();
  const [screenHeight, setScreenHeight] = useState(0);
  const [changePasswordModal, setChangePasswordModal] = useState(false);
  const openChangePasswordModal = () => setChangePasswordModal(true);
  const closeChangePasswordModal = () => setChangePasswordModal(false);

  useEffect(() => {
    localStorage.getItem("info") != null ? setLogin(true) : setLogin(false);
    setScreenHeight(window.innerHeight);
  }, [])

  const handleLogOut = () => {
    localStorage.removeItem("info")
    localStorage.removeItem("username")
    localStorage.removeItem("email")
    window.location.href = "/"
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
              <SubMenu className='subMenuSidebar' label={localStorage.getItem("username")} icon={<FaUserCircle/>}
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

            <ChangePasswordModal closeChangePasswordModal={closeChangePasswordModal} changePasswordModal={changePasswordModal} email={localStorage.getItem("email")} />
          </main>
        </HashRouter>
      </div>
    );
  }
}

export default App;
