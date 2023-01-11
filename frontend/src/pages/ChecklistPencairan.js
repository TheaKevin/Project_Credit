import React, { useEffect, useState } from 'react'
import axios from 'axios';
import Table from 'react-bootstrap/Table';
import Form from 'react-bootstrap/Form';

export const ChecklistPencairan = () => {
    const [transaction, setTransaction] = useState([])
    const [checked, setChecked] = useState([])
    useEffect(() => {
        getTransactionData()
    }, [])

    const getTransactionData = () => {
        axios.get("http://localhost:8080/").then( (res) => {
            setTransaction(res.data.data)
            console.log(res.data.message)
        })
    }

    const checkTransaction = (id) => {
        var validation = 0
        if (checked.length !== 0) {
            for (let index = 0; index < checked.length; index++) {
                if (checked[index].Custcode === transaction[id].Custcode) {
                    validation = 1
                    checked.splice(index, 1);
                    setChecked(checked => [...checked]);
                    break
                }
            }
        }

        if (validation === 0) {
            setChecked(checked => [...checked, {Custcode:transaction[id].Custcode}])
        }
    }

    const updateTransaction = () => {
        axios.patch('http://localhost:8080/updateTransaction', checked).then(() => getTransactionData())
    }
    
    return (
        <div>
            <Table striped bordered hover>
                <thead>
                    <tr>
                        <th>No.</th>
                        <th>PPK</th>
                        <th>Name</th>
                        <th>Channeling Company</th>
                        <th>Drawdown Date</th>
                        <th>Loan Amount</th>
                        <th>Loan Period</th>
                        <th>Interest Eff</th>
                        <th>Check</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        transaction ? (
                            transaction.map((value, key) => 
                                <tr key={key}>
                                    <td>{key+1}</td>
                                    <td>{value.PPK}</td>
                                    <td>{value.Name}</td>
                                    <td>{value.ChannelingCompany}</td>
                                    <td>{value.DrawdownDate}</td>
                                    <td>{value.LoanAmount}</td>
                                    <td>{value.LoanPeriod}</td>
                                    <td>{value.InterestEffective}</td>
                                    <td>
                                        <Form.Check onChange={() => checkTransaction(key)} />
                                    </td>
                                </tr>
                            )
                        ) : null
                    }
                </tbody>
            </Table>
            <button onClick={() => updateTransaction()}>Approve</button>
        </div>
    )
}
