import axios from 'axios';
import React, { useEffect, useState } from 'react'
import { Modal } from 'react-bootstrap';
import "../style/ChangePassword.css"

export const ChangePasswordModal = (props) => {
    const [password, setPassword] = useState();

    useEffect(() => {
        if (password) {
            const alertWrongInput = document.getElementById('alertWrongInput')

            axios.patch('http://localhost:8080/changePassword', password)
            .then(res => {
                if(res.data.message === "success"){
                    alertWrongInput.classList.add("d-none");
                    alert("Success");
                    setPassword(undefined)
                }
            })
            .catch(err => {
                document.getElementById("alertWrongInput").innerHTML = err.response.data.message;
                alertWrongInput.classList.remove("d-none");
            })
        }
    }, [password]);

    const handleChangePassword = (e) => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget)
        const alertWrongInput = document.getElementById('alertWrongInput')

        if (formData.get("oldPassword") === "") {
            document.getElementById("alertWrongInput").innerHTML = "Password lama harus diisi!";
            alertWrongInput.classList.remove("d-none");
        } else if (formData.get("newPassword") === "") {
            document.getElementById("alertWrongInput").innerHTML = "Password baru harus diisi!";
            alertWrongInput.classList.remove("d-none");
        } else if (formData.get("confirmPassword") === "") {
            document.getElementById("alertWrongInput").innerHTML = "Confirm Password harus diisi!";
            alertWrongInput.classList.remove("d-none");
        } else if (formData.get("newPassword") !== formData.get("confirmPassword")) {
            document.getElementById("alertWrongInput").innerHTML = "Confirm Password tidak sama!";
            alertWrongInput.classList.remove("d-none");
        } else if (formData.get("newPassword") === formData.get("oldPassword")) {
            document.getElementById("alertWrongInput").innerHTML = "Password lama tidak boleh sama dengan password baru!";
            alertWrongInput.classList.remove("d-none");
        } else {
            alertWrongInput.classList.add("d-none");
            setPassword({
                Email:props.email,
                OldPassword:formData.get("oldPassword"),
                NewPassword:formData.get("newPassword")
            })
        }
    }

    return (
        <Modal show={props.changePasswordModal} onHide={props.closeChangePasswordModal}
            aria-labelledby="contained-modal-title-vcenter"
            centered>
            <Modal.Header closeButton>
                <Modal.Title className="ms-auto">Change Password</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <div className="col-lg-12 d-flex justify-content-center">
                    <form onSubmit={(e) => handleChangePassword(e)} className="p-2 w-100">
                        <div id="alertWrongInput" className="alert alert-danger d-none" role="alert">
                            Confirm Password tidak sama!
                        </div>

                        <div className="input-group mb-4 mt-4">
                            <input type="password" className="input" name="oldPassword"/>
                            <label className="placeholderChangePassword">Old Password</label>    
                        </div>

                        <div className="input-group mb-4">
                            <input type="password" className="input" name="newPassword"/>
                            <label className="placeholderChangePassword">New Password</label>    
                        </div>

                        <div className="input-group mb-4">
                            <input type="password" className="input" name="confirmPassword"/>
                            <label className="placeholderChangePassword">Confirm Password</label>    
                        </div>

                        <div className="d-flex justify-content-around">
                            <button type="submit" className="buttonPrimary w-50">
                                Submit
                            </button>
                        </div>
                    </form>
                </div>
            </Modal.Body>
        </Modal>
    )
}
