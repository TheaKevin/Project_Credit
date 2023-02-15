import React, { useEffect, useState } from 'react'
import axios from 'axios';
import Table from 'react-bootstrap/Table';
import Form from 'react-bootstrap/Form';
import "../style/ChecklistPencairan.css"

export const ChecklistPencairan = () => {
    const [transaction, setTransaction] = useState([])
    const [company, setCompany] = useState([])
    const [branch, setBranch] = useState([])
    const [activeCompany, setActiveCompany] = useState()
    const [activeBranch, setActiveBranch] = useState()
    const [start, setStart] = useState()
    const [end, setEnd] = useState()
    const [checked, setChecked] = useState([])
    useEffect(() => {
        getTransactionData()
    }, [])

    const getTransactionData = () => {
        axios.get("http://localhost:8080/getTransaction").then( (res) => {
            setTransaction(res.data.data)
            setCompany(res.data.company)
            setBranch(res.data.branch)
            setActiveCompany(res.data.company[0].company_short_name)
            setActiveBranch(res.data.branch[0].code)
            setStart(new Date(0))
            setEnd(new Date())
            console.log(res.data.message)
        })
    }

    const checkTransaction = (id, e) => {
        if(e.target.checked) setChecked(checked => [...checked, {Custcode:transaction[id].Custcode}])
        else setChecked((i) => i.filter((j) => j.Custcode !== transaction[id].Custcode))
        // var validation = 0
        // if (checked.length !== 0) {
        //     for (let index = 0; index < checked.length; index++) {
        //         if (checked[index].Custcode === transaction[id].Custcode) {
        //             validation = 1
        //             checked.splice(index, 1);
        //             setChecked(checked => [...checked]);
        //             break
        //         }
        //     }
        // }

        // if (validation === 0) {
        //     setChecked(checked => [...checked, {Custcode:transaction[id].Custcode}])
        // }
    }

    const updateTransaction = () => {
        axios.patch('http://localhost:8080/updateTransaction', checked).then(() => getTransactionData())
    }

    const handleSubmit = (e) => {
        e.preventDefault()
        axios.get("http://localhost:8080/getTransactionFilter", {
            params: {
                branch: activeBranch,
                company: activeCompany,
                start: start,
                end: end
            }
        }).then( (res) => {
            setTransaction(res.data.data)
            console.log(res.data.message)
        })
    }
    
    return (
        <div className='d-flex flex-column w-100 px-5 pt-3'>
            <Form onSubmit={(e) => handleSubmit(e)} className="mb-4">
                <div className='groupFilter'>
                    <Form.Label className='branchLabel'>Branch</Form.Label>
                    <Form.Select className='branchSelect' value={activeBranch} onChange={(e) => setActiveBranch(e.target.value)} >
                        {
                            branch? (
                                branch.map((value, key) =>
                                    <option key={key} value={value.code}>{value.code} - {value.description}</option>
                                )
                            ) : null
                        }
                    </Form.Select>

                    <Form.Label className='companyLabel'>Company</Form.Label>
                    <Form.Select className='companySelect' value={activeCompany} onChange={(e) => setActiveCompany(e.target.value)} >
                        {
                            company? (
                                company.map((value, key) =>
                                    <option key={key}>{value.company_short_name}</option>
                                )
                            ) : null
                        }
                    </Form.Select>

                    <Form.Label className='startLabel' htmlFor='startDate'>Start</Form.Label>
                    <Form.Control className='startDate'
                        type='date'
                        id='startDate'
                        name='startDate'
                        onChange={(e) => setStart(new Date(e.target.value))} />

                    <Form.Label className='endLabel' htmlFor='endDate'>End</Form.Label>
                    <Form.Control className='endDate'
                        type='date'
                        id='endDate'
                        name='endDate'
                        onChange={(e) => setEnd(new Date(e.target.value))} />

                    <button className='resetButton' onClick={(e) => {
                        e.preventDefault()
                        getTransactionData()
                    }}>Reset</button>

                    <button className='submitButton' type='submit'>Submit</button>
                </div>
            </Form>
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
                                        <Form.Check onChange={(e) => checkTransaction(key, e)} />
                                    </td>
                                </tr>
                            )
                        ) : null
                    }
                </tbody>
            </Table>
            <button className='align-self-end' onClick={() => updateTransaction()}>Approve</button>
        </div>
    )
}
