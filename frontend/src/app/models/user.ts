import { RoleDTO } from './role';

export interface UserDTO {
    ID?: number;
    Username?: string;
    Role_id?: number;
    CreatedAt?: Date;
    UpdatedAt?: Date;
    DeletedAt?: Date;
    Role?: RoleDTO;
}
