uuid = require('uuid')
datetime = require('datetime')
-- Create a space --
box.schema.space.create('dialogs')

-- Specify field names and types --
box.space.dialogs:format({
    { name = 'id', type = 'uuid' },
    { name = 'from', type = 'uuid' },
    { name = 'to', type = 'uuid' },
    { name = 'text', type = 'string' },
    { name = 'timestamp', type = 'datetime' }
})

-- Create indexes --
box.space.dialogs:create_index('primary', { parts = { 'id' } })
box.space.dialogs:create_index('spatial', { type = 'tree', unique = false, parts = { 'from','to' } })
-- box.space.dialogs:create_index('band', { parts = { 'band_name' } })
-- box.space.dialogs:create_index('year_band', { parts = { { 'year' }, { 'band_name' } } })

-- Create a stored function --
box.schema.func.create('send', {
    body = [[
    function(from,to,text)
        return box.space.dialogs:insert({uuid.new(),uuid.fromstr(from),uuid.fromstr(to),text,datetime.now()})
    end
    ]]
})

-- box.schema.func.create('phello', {
--     body = [[
--     function(str)
--         return str
--     end
--     ]]
-- })


box.schema.func.create('list', {
    body = [[
    function(from,to)
        return box.space.dialogs.index.spatial:select({uuid.fromstr(from),uuid.fromstr(to)})
    end
    ]]
})