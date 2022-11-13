import React, { useEffect, useState } from 'react';

export const Calculate: React.FC = () => {

  const [price, setPrice] = useState<number>(0);
  const [debtor, setDebtor] = useState<string>("");
  const [lender, setLender] = useState<string>("");
  const [debts, setDebts] = useState<Array<any>>([]);
  const trackTextStyle: React.CSSProperties = {
    fontFamily: "revert-layer",
    fontSize: "2.8vw",
  }

  const getDebts = () => {
    const url = 'debts';
    const method = 'GET';
    const headers = {
      'Content-Type': 'application/json; charset=UTF-8',
    };

    fetch(url, {method, headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
    }).then((json) => {
      setDebts(json);
    });
  }
  const calculate = () => {
    const url = 'calculate';
    const method = 'GET';
    const headers = {
      'Content-Type': 'application/json; charset=UTF-8',
    };

    fetch(url, {method, headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
    }).then((json) => {
      setDebtor(json.debtor);
      setLender(json.lender);
      setPrice(json.price)
    });
  }

  const settle = () => {
    const url = 'settle';
    const method = 'PUT';
    const headers = {
      'Content-Type': 'application/json; charset=UTF-8',
    };

    fetch(url, {method, headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
    }).then((json) => {
      calculate();
      getDebts();
    });
  }

  useEffect(() => {
    console.log("component is rendered!!!");
    calculate();
    getDebts();
  }, []);

  return (
    <>
      <div style={trackTextStyle}>{debtor} が {lender} に {price} 円の借金</div>
      <button onClick={settle}>精算する</button>
      {debts.map((value, index) => {
        return <li key={index}>{value.debtor} が {value.lender} に {value.price} 円の借金：{value.memo}</li>
      })}
    </>
  );
}