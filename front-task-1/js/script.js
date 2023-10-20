const tableBody = document.getElementById('currencyTable');

fetch('https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1')
    .then(response => response.json())
    .then(data => {
        data.forEach(currency => {
            const { id, symbol, name } = currency;

            const row = document.createElement('tr');
            row.dataset.symbol = symbol;

            const idCell = document.createElement('td');
            idCell.textContent = id;

            const symbolCell = document.createElement('td');
            symbolCell.textContent = symbol;

            const nameCell = document.createElement('td');
            nameCell.textContent = name;

            row.appendChild(idCell);
            row.appendChild(symbolCell);
            row.appendChild(nameCell);

            tableBody.appendChild(row);
        });
    })

    