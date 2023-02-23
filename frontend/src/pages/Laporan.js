import React, { useEffect, useState } from 'react'
import axios from 'axios';
import Table from 'react-bootstrap/Table';
import Form from 'react-bootstrap/Form';
import "../style/Laporan.css"

export const Laporan = () => {
    const [report, setReport] = useState([])
    const [company, setCompany] = useState([])
    const [branch, setBranch] = useState([])
    const [activeCompany, setActiveCompany] = useState()
    const [activeBranch, setActiveBranch] = useState()
    const [start, setStart] = useState()
    const [end, setEnd] = useState()
    const [status, setStatus] = useState()

    useEffect(() => {
        getReportData()
    }, [])

    const getReportData = () => {
        axios.get("http://localhost:8080/getReport").then( (res) => {
            setReport(res.data.data)
            setCompany(res.data.company)
            setBranch(res.data.branch)
            setActiveCompany(res.data.company[0].company_short_name)
            setActiveBranch(res.data.branch[0].code)
            setStart(new Date(0))
            setEnd(new Date())
            setStatus(0)
            console.log(res.data.message)
        })
    }

    const handleSubmit = (e) => {
        e.preventDefault()
        axios.get("http://localhost:8080/getReportFilter", {
            params: {
                branch: activeBranch,
                company: activeCompany,
                start: start,
                end: end,
                status: status
            }
        }).then( (res) => {
            setReport(res.data.data)
            console.log(res.data.message)
        })
    }
    
    return (
        <div className='d-flex flex-column w-100 px-5 pt-3'>
            <Form onSubmit={(e) => handleSubmit(e)} className="mb-4">
                <div className='groupFilterLaporan'>
                    <Form.Group className='d-flex flex-row align-items-center branchLaporan'>
                        <Form.Label className='me-2'>Branch</Form.Label>
                        <Form.Select value={activeBranch} onChange={(e) => setActiveBranch(e.target.value)} >
                            {
                                branch? (
                                    branch.map((value, key) =>
                                        <option key={key} value={value.code}>{value.code} - {value.description}</option>
                                    )
                                ) : null
                            }
                        </Form.Select>
                    </Form.Group>

                    <Form.Group className='d-flex flex-row align-items-center companyLaporan'>
                        <Form.Label className='me-2'>Company</Form.Label>
                        <Form.Select value={activeCompany} onChange={(e) => setActiveCompany(e.target.value)} >
                            {
                                company? (
                                    company.map((value, key) =>
                                        <option key={key}>{value.company_short_name}</option>
                                    )
                                ) : null
                            }
                        </Form.Select>
                    </Form.Group>

                    <Form.Group className='d-flex flex-row align-items-center approveLaporan'>
                        <Form.Label className='me-2 w-50'>Approval Status</Form.Label>
                        <Form.Select className='w-50' value={status} onChange={(e) => setStatus(e.target.value)} >
                            <option value={0}>0</option>
                            <option value={1}>1</option>
                        </Form.Select>
                    </Form.Group>

                    <Form.Group className='d-flex flex-row align-items-center startLaporan'>
                        <Form.Label className='me-2' htmlFor='startDate'>Start</Form.Label>
                        <Form.Control
                            type='date'
                            id='startDate'
                            name='startDate'
                            onChange={(e) => setStart(new Date(e.target.value))} />
                    </Form.Group>

                    <Form.Group className='d-flex flex-row align-items-center endLaporan'>
                        <Form.Label className='me-2' htmlFor='endDate'>End</Form.Label>
                        <Form.Control
                            type='date'
                            id='endDate'
                            name='endDate'
                            onChange={(e) => setEnd(new Date(e.target.value))} />
                    </Form.Group>
                    
                    <button className='resetButtonLaporan buttonPrimary' onClick={(e) => {
                        e.preventDefault()
                        getReportData()
                    }}>All</button>

                    <button className='submitButtonLaporan buttonPrimary' type='submit'>Submit</button>
                </div>
            </Form>
            <Table striped bordered hover responsive>
                <thead className='tableHeadWrapper'>
                    <tr>
                        <th>No.</th>
                        <th>PPK</th>
                        <th>Name</th>
                        <th>Channeling Company</th>
                        <th>Drawdown Date</th>
                        <th>Loan Amount</th>
                        <th>Loan Period</th>
                        <th>Interest Eff</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        report ? (
                            report.map((value, key) => 
                                <tr key={key}>
                                    <td>{key+1}</td>
                                    <td>{value.PPK}</td>
                                    <td>{value.Name}</td>
                                    <td>{value.ChannelingCompany}</td>
                                    <td>{value.DrawdownDate}</td>
                                    <td>{value.LoanAmount}</td>
                                    <td>{value.LoanPeriod}</td>
                                    <td>{value.InterestEffective}</td>
                                </tr>
                            )
                        ) : null
                    }
                </tbody>
            </Table>
        </div>
    )
}
