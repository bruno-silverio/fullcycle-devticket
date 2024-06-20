// Classes que vao modelar os dados que estao sendo recebidos pela requisicao HTTP
export class CriarEventoRequest {
  nome: string;
  descricao: string;
  data: string;
  preco: number;
}
