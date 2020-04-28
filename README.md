1) Você deverá criar uma simples aplicação que possua uma função chamada soma que receberá dois parâmetros e retornará a adição desses dois valores.

Essa função deverá ser chamada na função main do programa. Quando executada, deverá exibir da na tela o resultado da soma de 5 + 5.

2) Crie um teste unitário para essa função.

3) Ative um processo de CI que execute os seguintes passos:

Executar o teste unitário
Push da imagem gerada no processo de CI no Container Registry do GCP
Ative a App do Cloud Build no Github para que cada pull request execute automaticamente o processo de CI
Dicas: 

Para fazer o push da imagem, reveja o vídeo onde fazemos a instalação do docker-compose.
Para executar os comandos Go no Cloud Build, lembre-se que o GCP fornece diversas imagens prontas para isso em seu registro público "cloud-builders"
No cloudbuild.yaml provavelmente você deverá setar uma variável de ambiente informando o GOPATH em cada etapa. Recomendo que verifique a documentação do Cloud Build na sessão da linguagem Go.
Crie uma pasta src/[nome-de-seu-projeto-ou-pacote] e coloquei seus arquivos .go nessa pasta.