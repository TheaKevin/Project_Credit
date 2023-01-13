import './App.css';
import { useState } from 'react';
import { Route, Switch, Redirect, Routes, HashRouter, Link } from 'react-router-dom';
import { FaBars, FaMoneyBill, FaRegClipboard, FaUserCircle } from 'react-icons/fa';
import { Sidebar, Menu, MenuItem, SubMenu, useProSidebar } from 'react-pro-sidebar';
import 'bootstrap/dist/css/bootstrap.min.css'
import sidebarBg from './assets/bg1.jpg'
import { ChecklistPencairan } from './pages/ChecklistPencairan';
import { Laporan } from './pages/Laporan';

function App() {
  const { collapseSidebar } = useProSidebar();

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
            <SubMenu label="User" icon={<FaUserCircle/>}>
              <MenuItem> Change Password </MenuItem>
              <MenuItem> Logout </MenuItem>
            </SubMenu>
          </Menu>
        </Sidebar>
        <main className='w-100 mx-5 my-3'>
          <Routes>
            <Route path="/" element={ <ChecklistPencairan  /> }/>
            <Route path="/Laporan" element={ <Laporan  /> }/>
          </Routes>
        </main>
      </HashRouter>
    </div>
  );
}

export default App;
