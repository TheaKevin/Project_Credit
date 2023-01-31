import './App.css';
import { useEffect, useState } from 'react';
import { Route, Switch, Redirect, Routes, HashRouter, Link, Navigate } from 'react-router-dom';
import { FaBars, FaMoneyBill, FaRegClipboard, FaUserCircle } from 'react-icons/fa';
import { Sidebar, Menu, MenuItem, SubMenu, useProSidebar } from 'react-pro-sidebar';
import 'bootstrap/dist/css/bootstrap.min.css'
import sidebarBg from './assets/bg1.jpg'
import { ChecklistPencairan } from './pages/ChecklistPencairan';
import { Laporan } from './pages/Laporan';
import { Login } from './pages/Login';
import { ChangePasswordModal } from './pages/ChangePasswordModal';

function App() {
  const { collapseSidebar } = useProSidebar();
  const [isLoggedIn, setLogin] = useState();
  const [changePasswordModal, setChangePasswordModal] = useState(false);
  const openChangePasswordModal = () => setChangePasswordModal(true);
  const closeChangePasswordModal = () => setChangePasswordModal(false);

  useEffect(() => {
    localStorage.getItem("info") != null ? setLogin(true) : setLogin(false)
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
      <div style={{ display: 'flex' }}>
        <HashRouter>
          <Sidebar image={sidebarBg}>
            <Menu>
              <MenuItem icon={<FaBars/>} onClick={() => collapseSidebar()} />
              <SubMenu label="Transaksi" icon={<FaMoneyBill />}>
                <MenuItem routerLink={<Link to="/" />}> Checklist Pencairan </MenuItem>
              </SubMenu>
              <MenuItem icon={<FaRegClipboard/>} routerLink={<Link to="/Laporan" />}> Laporan </MenuItem>
              <SubMenu label={localStorage.getItem("username")} icon={<FaUserCircle/>}>
                <MenuItem onClick={() => openChangePasswordModal()}> Change Password </MenuItem>
                <MenuItem onClick={() => handleLogOut()}> Logout </MenuItem>
              </SubMenu>
            </Menu>
          </Sidebar>
          <main className='w-100 mx-5 my-3'>
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
