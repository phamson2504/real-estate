import React, { useState, useEffect } from 'react'
import './RequestedProperties.css'
import useAuth from '../../hooks/useAuth';

export default function RequestedProperties() {
    const { user, authAxios } = useAuth()
    const [transaction, setTransaction] = useState([]);
    const fetchData = async () => {
        try {
            const response = await authAxios.get(`/transaction/get-transaction-offered-seller?id=${user.id}`);
            setTransaction(response.data.transactions)
            console.log(response.data.transactions)
        } catch (error) {
            console.error("Error fetching data:", error);
        }
    };

    useEffect(() => {
        fetchData();
    }, [])

    const handleAction = async (transactionId, status) => {
        try {
            const response = await authAxios.post(
                '/transaction/request-transaction-for-seller',
                {
                    id: transactionId,
                    status: status,
                },
            );
            console.log('API Response:', response.data);

            var statusChange
            if (status === 1) 
                statusChange = "accepted"
            else
                statusChange = "reject"

            setTransaction((prevTransactions) =>
                prevTransactions.map((transaction) =>
                  transaction.id === transactionId
                    ? { ...transaction, status: statusChange }
                    : transaction
                )
              );
              alert(`Update successfully`);
        } catch (error) {
            console.error('API Error:', error);
            alert(`Failed to transaction`);
        }
    };

    return (
        <div className="pr-container">
            <h2>Requested Properties</h2>
            <table className="rp-table">
                <thead className="rp-table-header">
                    <tr>
                        <th className="rp-table-title">Property Title</th>
                        <th className="rp-table-location">Property Location</th>
                        <th className="rp-table-buyer-name">Buyer Name</th>
                        <th className="rp-table-buyer-email">Buyer Email</th>
                        <th className="rp-table-price">Offered Price</th>
                        <th className="rp-table-actions">Status Action</th>
                    </tr>
                </thead>
                <tbody className="rp-table-body">
                    {transaction && transaction.map(t => (
                        <tr key={t.id}>
                            <td className="rp-table-title">{t.Properties.title}</td>
                            <td className="rp-table-location">{t.Properties.location}/</td>
                            <td className="rp-table-buyer-name">{t.Properties.agent.name}</td>
                            <td className="rp-table-buyer-email">{t.Properties.agent.email}</td>
                            <td className="rp-table-price">${t.amount}</td>
                            <td className="rp-table-actions">
                                {t.status === "spending" ? (
                                    <>
                                        <button className="rp-accept-btn" onClick={() => handleAction(t.id, 1)}>Accept</button>
                                        <button className="rp-reject-btn" onClick={() => handleAction(t.id, 3)}>Reject</button>
                                    </>
                                ) : t.status === "accepted" ? (
                                    <>
                                        <span>Accept</span>
                                    </>
                                ) : (
                                    <>
                                        <span>Reject</span>
                                    </>
                                )}
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    )
}
