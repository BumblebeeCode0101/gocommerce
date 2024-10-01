# GoCommerce

GoCommerce is a free and open-source eCommerce platform designed for scalability and customization.

## Installation
To install GoCommerce, follow these steps:

1. **Clone the Repository:**
   
   ```bash
   git clone https://github.com/BumblebeeCode0101/gocommerce.git && cd gocommerce
   
2. **Set Up Environment Variables**

   ```bash
   cp .env.example .env

3. **Install Dependencies**
   ```bash
   go mod tidy

   pip3 install -r requirements.txt

4. **Generate App Key**
   ```bash
   python3 tools/generate_app_key.py

5. **Build GoCommerce**
   ```bash
   go build src/main.go -o gocommerce

6. **Run GoCommerce**
   ```bash
   ./gocommerce
