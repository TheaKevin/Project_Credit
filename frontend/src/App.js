import './App.css';
import { useState } from 'react';
import { Route, Switch, Redirect, Routes, HashRouter } from 'react-router-dom';
import { FaBars, FaMoneyBill, FaRegClipboard, FaUserCircle } from 'react-icons/fa';
import { Sidebar, Menu, MenuItem, SubMenu, useProSidebar } from 'react-pro-sidebar';
import 'bootstrap/dist/css/bootstrap.min.css'
import sidebarBg from './assets/bg1.jpg'
import { ChecklistPencairan } from './pages/ChecklistPencairan';

function App() {
  const { collapseSidebar } = useProSidebar();

  return (
    <div style={{ display: 'flex', height: '100%' }}>
      <HashRouter>
        <Sidebar image={sidebarBg}>
          <Menu>
            <MenuItem icon={<FaBars/>} onClick={() => collapseSidebar()} />
            <SubMenu label="Transaksi" icon={<FaMoneyBill />}>
              <MenuItem> Checklist Pencairan </MenuItem>
            </SubMenu>
            <MenuItem icon={<FaRegClipboard/>}> Laporan </MenuItem>
            <SubMenu label="User" icon={<FaUserCircle/>}>
              <MenuItem> Change Password </MenuItem>
              <MenuItem> Logout </MenuItem>
            </SubMenu>
          </Menu>
        </Sidebar>
        <main>
          <Routes>
            <Route path="/" element={ <ChecklistPencairan  /> }/>
          </Routes>
        </main>
      </HashRouter>
    </div>
  );
}

export default App;
