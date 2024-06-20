// Classes que vao modelar os dados que estao sendo recebidos pela requisicao HTTP
export class CreateEventDto {
  name: string;
  description: string;
  date: string;
  price: number;
}
