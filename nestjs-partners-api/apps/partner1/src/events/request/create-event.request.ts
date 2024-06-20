// Classes que vao modelar os dados que estao sendo recebidos pela requisicao HTTP
export class CreateEventRequest {
  name: string;
  description: string;
  date: string;
  price: number;
}
