import React from 'react'
import "../style/Loading.css"

export const Loading = () => {
    return (
        <div className='d-flex justify-content-center'>
            <div className='lds-ring'><div></div><div></div><div></div><div></div></div>
        </div>
    )
}
