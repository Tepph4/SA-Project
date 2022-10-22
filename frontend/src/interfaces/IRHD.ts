import { RoomInterface } from "./IRoom";
import { DeviceInterface } from "./IDevice";
import { UsersInterface } from "./IUser";

export interface RHDsInterface {
    ID?: number; 

    RoomID?: number;
    Room?: RoomInterface;

    DeviceID?: number;
    Device?: DeviceInterface;
    
    UserID?: number;
    User?: UsersInterface;
    
  }