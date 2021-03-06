/* tslint:disable */
/* eslint-disable */
/**
 * SUT SE Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntFurnitureDetailEdges,
    EntFurnitureDetailEdgesFromJSON,
    EntFurnitureDetailEdgesFromJSONTyped,
    EntFurnitureDetailEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntFurnitureDetail
 */
export interface EntFurnitureDetail {
    /**
     * DateAdd holds the value of the "date_add" field.
     * @type {string}
     * @memberof EntFurnitureDetail
     */
    dateAdd?: string;
    /**
     * 
     * @type {EntFurnitureDetailEdges}
     * @memberof EntFurnitureDetail
     */
    edges?: EntFurnitureDetailEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntFurnitureDetail
     */
    id?: number;
}

export function EntFurnitureDetailFromJSON(json: any): EntFurnitureDetail {
    return EntFurnitureDetailFromJSONTyped(json, false);
}

export function EntFurnitureDetailFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntFurnitureDetail {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'dateAdd': !exists(json, 'date_add') ? undefined : json['date_add'],
        'edges': !exists(json, 'edges') ? undefined : EntFurnitureDetailEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntFurnitureDetailToJSON(value?: EntFurnitureDetail | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'date_add': value.dateAdd,
        'edges': EntFurnitureDetailEdgesToJSON(value.edges),
        'id': value.id,
    };
}


