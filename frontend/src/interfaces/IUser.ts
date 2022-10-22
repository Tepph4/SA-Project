import { RolesInterface } from "./IRole";

export interface UsersInterface{
    ID? : number,
    Name? : string,
    Email? : string,
    Password? : string;
    RoleID?: string;
    Role?: RolesInterface;
}